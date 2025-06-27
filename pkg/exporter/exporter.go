package exporter

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/Jeffail/gabs/v2"
	"github.com/Tnze/go-mc/nbt"
	mcnet "github.com/Tnze/go-mc/net"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace                  = "minecraft"
	Forge                      = "forge"
	PaperMC                    = "papermc"
	Fabric                     = "fabric"
	PurpurMC                   = "purpurmc"
	rconListCommand            = "list"
	rconForgeTpsCommand        = "forge tps"
	rconForgeEntityListCommand = "forge entity list"
	rconTpsCommand             = "tps"
	rconFabricTpsCommand       = "fabric tps"
)

// See for all details on the statistics of Minecraft https://minecraft.fandom.com/wiki/Statistics

type Exporter struct {
	rcon               *RCON
	logger             *slog.Logger
	world              string
	source             string
	serverStats        string
	disabledMetrics    map[string]bool
	playerOnlineRegexp *regexp.Regexp
	overallRegexp      *regexp.Regexp
	dimensionRegexp    *regexp.Regexp
	entityListRegexp   *regexp.Regexp
	paperMcTpsRegexp   *regexp.Regexp
	purpurMcTpsRegexp  *regexp.Regexp

	// via advancements
	// playerAdvancements *prometheus.Desc

	// via RCON
	playerOnline    *prometheus.Desc
	dimTps          *prometheus.Desc
	dimTicktime     *prometheus.Desc
	overallTps      *prometheus.Desc
	overallTicktime *prometheus.Desc
	entities        *prometheus.Desc
	tpsPaperMC      *prometheus.Desc

	// via stats
	playerStat *prometheus.Desc

	blocksMined            *prometheus.Desc
	entitiesKilled         *prometheus.Desc
	playerKilledBy         *prometheus.Desc
	item                   *prometheus.Desc
	animalsBred            *prometheus.Desc
	cleanArmor             *prometheus.Desc
	cleanBanner            *prometheus.Desc
	openBarrel             *prometheus.Desc
	bellRing               *prometheus.Desc
	eatCakeSlice           *prometheus.Desc
	fillCauldron           *prometheus.Desc
	openChest              *prometheus.Desc
	damageDealt            *prometheus.Desc
	damageReceived         *prometheus.Desc
	inspected              *prometheus.Desc
	minecraftMovement      *prometheus.Desc
	openEnderChest         *prometheus.Desc
	fishCaught             *prometheus.Desc
	leaveGame              *prometheus.Desc
	interaction            *prometheus.Desc
	itemsDropped           *prometheus.Desc
	itemsEntchanted        *prometheus.Desc
	jump                   *prometheus.Desc
	mobKills               *prometheus.Desc
	musicDiscsPlayed       *prometheus.Desc
	noteBlocksPlayed       *prometheus.Desc
	noteBlocksTuned        *prometheus.Desc
	numberOfDeaths         *prometheus.Desc
	plantsPotted           *prometheus.Desc
	playerKills            *prometheus.Desc
	raidsTriggered         *prometheus.Desc
	raidsWon               *prometheus.Desc
	shulkerBoxCleaned      *prometheus.Desc
	shulkerBoxesOpened     *prometheus.Desc
	sneakTime              *prometheus.Desc
	talkedToVillager       *prometheus.Desc
	targetsHit             *prometheus.Desc
	timePlayed             *prometheus.Desc
	timeSinceDeath         *prometheus.Desc
	timeSinceLastRest      *prometheus.Desc
	timesWorldOpen         *prometheus.Desc
	timesSleptInBed        *prometheus.Desc
	tradedWithVillagers    *prometheus.Desc
	trappedChestsTriggered *prometheus.Desc
	waterTakenFromCauldron *prometheus.Desc
}

// PlayerDB represents the structure of a player
type PlayerDB struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
	Success bool   `json:"success"`
}
type Meta struct {
	CachedAt int `json:"cached_at"`
}
type Properties struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}
type Player struct {
	Meta        Meta         `json:"meta"`
	Username    string       `json:"username"`
	ID          string       `json:"id"`
	RawID       string       `json:"raw_id"`
	Avatar      string       `json:"avatar"`
	SkinTexture string       `json:"skin_texture"`
	Properties  []Properties `json:"properties"`
	NameHistory []any        `json:"name_history"`
}
type Data struct {
	Player Player `json:"player"`
}

type PlayerData struct {
	XpLevel   int32
	XpTotal   int32
	Score     int32
	Health    interface{}
	FoodLevel int32 `nbt:"foodLevel"`
	Bukkit    struct {
		LastKnownName string `nbt:"lastKnownName"`
	} `nbt:"bukkit"`
}

type RCON struct {
	rconClient   mcnet.RCONClientConn
	rconServer   string
	rconPassword string
}

func createRCONClient(server, password string, logger *slog.Logger) *RCON {
	var rconClient mcnet.RCONClientConn
	if len(password) > 0 {
		var err error
		rconClient, err = mcnet.DialRCON(server, password)
		if err != nil {
			logger.Error("failed to connect to RCON", "err", err)
			rconClient = nil
		}
	}
	return &RCON{
		rconClient:   rconClient,
		rconServer:   server,
		rconPassword: password,
	}
}

func New(server, password, world, source, serverStats string, disabledMetrics map[string]bool, logger *slog.Logger) (*Exporter, error) {
	rcon := createRCONClient(server, password, logger)
	return &Exporter{
		rcon:               rcon,
		logger:             logger,
		world:              world,
		source:             source,
		serverStats:        serverStats,
		playerOnlineRegexp: regexp.MustCompile(":(.*)"),
		overallRegexp:      regexp.MustCompile(`Overall\s*:\sMean tick time:\s(\d*.\d*)\sms\.\sMean\sTPS:\s(\d*.\d*)`),
		dimensionRegexp:    regexp.MustCompile(`Dim\s(.*):(.*)\s\(.*\):\sMean tick time:\s(\d*.\d*)\sms\.\sMean\sTPS:\s(\d*.\d*)`),
		entityListRegexp:   regexp.MustCompile(`(\d+):\s(.*):(.*)`),
		paperMcTpsRegexp:   regexp.MustCompile(`§.TPS from last\s1m,\s5m,\s15m:\s§.(\d.*),\s§.(\d.*),\s§.(\d.*)`),
		purpurMcTpsRegexp:  regexp.MustCompile(`§.TPS from last\s5s,\s1m,\s5m,\s15m:\s§.(\d.*),\s§.(\d.*),\s§.(\d.*),\s§.(\d.*)`),
		disabledMetrics:    disabledMetrics,
		playerOnline: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "player_online_total"),
			"Players currently online (1 if player is online)",
			[]string{"player"},
			nil,
		),
		playerStat: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_stat_total"),
			"Statistic related to the player: xp, current_xp, food_level, health, score, advancements",
			[]string{"player", "stat"},
			nil,
		),
		blocksMined: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "blocks_mined_total"),
			"Statistic related to the number of blocks a player mined",
			[]string{"player", "namespace", "block"},
			nil,
		),
		entitiesKilled: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "entities_killed_total"),
			"Statistics related to the number of entities a player killed",
			[]string{"player", "namespace", "entity"},
			nil,
		),
		playerKilledBy: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "killed_by_total"),
			"Statistics related to the number of times a player was killed by entities",
			[]string{"player", "namespace", "entity"},
			nil,
		),
		item: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "item_actions_total"),
			"Statistics related to items and the number of times they were used, picked up, dropped or broken",
			[]string{"player", "namespace", "entity", "action"},
			nil,
		),
		animalsBred: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "animals_breded_total"),
			"The number of times the player bred two mobs",
			[]string{"player"},
			nil,
		),
		cleanArmor: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "cleaned_armors_total"),
			"The number of times the player washed dyed leather armor with a cauldron",
			[]string{"player"},
			nil,
		),
		//------
		cleanBanner: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "cleaned_banner_total"),
			"The number of times the player washed a banner with a cauldron",
			[]string{"player"},
			nil,
		),
		openBarrel: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "opened_barrels_total"),
			"The number of times the player opened a barrel",
			[]string{"player"},
			nil,
		),
		bellRing: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "bells_ringed_total"),
			"The number of times the player rang a bell",
			[]string{"player"},
			nil,
		),
		eatCakeSlice: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "cake_slices_eaten_total"),
			"The number of times the player ate cake",
			[]string{"player"},
			nil,
		),
		fillCauldron: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "filled_cauldrons_total"),
			"The number of times the player filled a cauldron with a water bucket",
			[]string{"player"},
			nil,
		),
		openChest: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "opened_chests_total"),
			"The number of times the player opened a chest",
			[]string{"player"},
			nil,
		),
		damageDealt: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_dealt_total"),
			"The amount of damage the player has dealt of different types (in tenths of 1♥)",
			[]string{"player", "type"},
			nil,
		),

		damageReceived: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_received_total"),
			"The amount of damage the player has taken of different types (in tenths of 1♥)",
			[]string{"player", "type"},
			nil,
		),

		minecraftMovement: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "movement_meters_total"),
			"The total distance the player traveled by different methods (ladders, boats, swimming etc.)",
			[]string{"player", "means"},
			nil,
		),
		inspected: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "inspected_total"),
			"The number of times the player inspected a dispenser, hopper or dropper",
			[]string{"player", "entity"},
			nil,
		),
		openEnderChest: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "opened_enderchests_total"),
			"The number of times the player opened an ender chest",
			[]string{"player"},
			nil,
		),
		fishCaught: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "fishs_caught_total"),
			"The number of times the player caught fish",
			[]string{"player"},
			nil,
		),
		leaveGame: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "games_left_total"),
			"The number of times the player clicked \"Save and quit to title\"",
			[]string{"player"},
			nil,
		),
		interaction: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interactions_total"),
			"The number of times the player interacted with different entities",
			[]string{"player", "entity"},
			nil,
		),
		itemsDropped: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "items_dropped_total"),
			"The number of items the player dropped",
			[]string{"player"},
			nil,
		),
		itemsEntchanted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "items_enchanted_total"),
			"The number of items the player enchanted",
			[]string{"player"},
			nil,
		),
		jump: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "jumps_total"),
			"The number of times the player jumped",
			[]string{"player"},
			nil,
		),
		mobKills: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "mobs_killed_total"),
			"The number of mobs the player killed",
			[]string{"player"},
			nil,
		),
		musicDiscsPlayed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "records_played_total"),
			"The number of times the player played a music disc on a jukebox",
			[]string{"player"},
			nil,
		),
		noteBlocksPlayed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "noteblocks_played_total"),
			"The number of times the player hit a note block",
			[]string{"player"},
			nil,
		),
		noteBlocksTuned: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "noteblocks_tuned_total"),
			"The number of times the player tuned a note block",
			[]string{"player"},
			nil,
		),
		numberOfDeaths: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "deaths_total"),
			"The number of times the player died",
			[]string{"player"},
			nil,
		),
		plantsPotted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "pots_flowered_total"),
			"The number of times the player planted a plant in a flower pot",
			[]string{"player"},
			nil,
		),
		playerKills: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "players_killed_total"),
			"The number of times the player killed a player",
			[]string{"player"},
			nil,
		),
		raidsTriggered: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "triggered_raids_total"),
			"The number of times the player triggered a raid",
			[]string{"player"},
			nil,
		),
		raidsWon: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "won_raids_total"),
			"The number of times the player won a raid",
			[]string{"player"},
			nil,
		),
		shulkerBoxCleaned: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "shulker_boxes_cleaned_total"),
			"The number of times the player washed a shulker box with a cauldron",
			[]string{"player"},
			nil,
		),
		shulkerBoxesOpened: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "shulker_boxes_opened_total"),
			"The number of times the player opened a shulker box",
			[]string{"player"},
			nil,
		),
		sneakTime: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "sneak_time_ticks_total"),
			"The number of ticks the player has spent sneaking",
			[]string{"player"},
			nil,
		),
		talkedToVillager: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "talked_to_villagers_total"),
			"The number of times the player spoke with a villager (opened the trading GUI)",
			[]string{"player"},
			nil,
		),
		targetsHit: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "targets_hit_total"),
			"The number of times the player shot a target block",
			[]string{"player"},
			nil,
		),
		timePlayed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "play_time_ticks_total"),
			"The number of ticks the player has played",
			[]string{"player"},
			nil,
		),
		timeSinceDeath: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "time_since_death_ticks_total"),
			"The number of ticks since the player's last death",
			[]string{"player"},
			nil,
		),
		timeSinceLastRest: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "time_since_rest_ticks_total"),
			"The number of ticks since the player's last rest (this is used to spawn phantoms)",
			[]string{"player"},
			nil,
		),
		timesWorldOpen: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "total_world_time_ticks_total"),
			"The number of ticks the player has been in the world",
			[]string{"player"},
			nil,
		),
		timesSleptInBed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "sleep_in_bed_ticks_total"),
			"The number of times the player slept in a bed",
			[]string{"player"},
			nil,
		),
		tradedWithVillagers: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "traded_with_villagers_total"),
			"The number of times the player traded with a villager",
			[]string{"player"},
			nil,
		),
		trappedChestsTriggered: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "triggered_trapped_chests_total"),
			"The number of times the player opened a trapped chest",
			[]string{"player"},
			nil,
		),
		waterTakenFromCauldron: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "used_cauldrons_total"),
			"The number of times the player took water from cauldrons with glass bottles",
			[]string{"player"},
			nil,
		),

		dimTps: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "dimension_tps_total"),
			"The number of ticks per second in a certain dimension",
			[]string{"namespace", "dimension"},
			nil,
		),

		dimTicktime: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "dimension_ticktime_total"),
			"The mean tick time in a certain dimension",
			[]string{"namespace", "dimension"},
			nil,
		),
		overallTps: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "tps_total"),
			"The overall mean ticks per second in the server",
			[]string{},
			nil,
		),
		overallTicktime: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "ticktime_total"),
			"The overall mean tick time in the server",
			[]string{},
			nil,
		),
		entities: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "active_entity_total"),
			"The number and type of an active entity on the server",
			[]string{"namespace", "entity"},
			nil,
		),

		tpsPaperMC: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "tps_total_bucket"),
			"The number of ticks per second in PaperMC",
			[]string{},
			nil,
		),
	}, nil
}

func (e *Exporter) getPlayerStats(ch chan<- prometheus.Metric) error {
	files, err := os.ReadDir(e.world + "/playerdata")
	if err != nil {
		return err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".dat" && !strings.Contains(file.Name(), "_cyclic") {
			id := strings.TrimSuffix(file.Name(), ".dat")
			f, err := os.Open(e.world + "/playerdata/" + file.Name())
			if err != nil {
				return err
			}

			r, err := gzip.NewReader(f)
			if err != nil {
				return err
			}

			var data PlayerData
			_, err = nbt.NewDecoder(r).Decode(&data)
			if err != nil {
				return err
			}
			err = r.Close()
			if err != nil {
				return err
			}
			err = f.Close()
			if err != nil {
				return err
			}

			var player PlayerDB
			switch e.source {
			case "mojang":
				req, err := http.NewRequest("GET", fmt.Sprintf("https://playerdb.co/api/player/minecraft/%s", id), nil)
				if err != nil {
					e.logger.Error("Failed to create request", "err", err)
					return err
				}
				req.Header.Set("User-Agent", "github.com/dirien/minecraft-prometheus-exporter")
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					e.logger.Error("Failed to connect to playerdb.co", "err", err)
					return err
				}

				if resp.StatusCode == 200 {
					if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
						e.logger.Error("Failed to connect decode response", "err", err)
						return err
					}
				} else {
					return fmt.Errorf("error retrieving player info from playerdb.co: %s", fmt.Sprintf("Status Code: %d", resp.StatusCode))
				}

				err = resp.Body.Close()
				if err != nil {
					return err
				}
			case "bukkit":
				if data.Bukkit.LastKnownName == "" {
					return fmt.Errorf("error retrieving player info from nbt: bukkit name is unknown")
				}

				player = PlayerDB{
					Data: Data{
						Player: Player{
							RawID:    id,
							Username: data.Bukkit.LastKnownName,
						},
					},
				}
			default:
				player = PlayerDB{
					Data: Data{
						Player: Player{
							RawID:    id,
							Username: id,
						},
					},
				}
			}

			ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.XpTotal), player.Data.Player.Username, "xp")
			ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.XpLevel), player.Data.Player.Username, "current_xp")
			ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.Score), player.Data.Player.Username, "score")
			ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.FoodLevel), player.Data.Player.Username, "food_level")
			health, err := strconv.ParseFloat(fmt.Sprint(data.Health), 64)
			if err != nil {
				return err
			}
			ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, health, player.Data.Player.Username, "health")

			err = e.advancements(id, ch, player.Data.Player.Username)
			if err != nil {
				return err
			}

			byteValue, err := os.ReadFile(e.world + "/stats/" + id + ".json")
			if err != nil {
				e.logger.Error(fmt.Sprintf("Stats file for player %s not exist", player.Data.Player.Username))
			} else {
				jsonParsed, err := gabs.ParseJSON(byteValue)
				if err != nil {
					return err
				}

				e.playerStats(jsonParsed, e.blocksMined, "minecraft:mined", ch, player.Data.Player.Username, "")
				e.playerStats(jsonParsed, e.entitiesKilled, "minecraft:killed", ch, player.Data.Player.Username, "")
				e.playerStats(jsonParsed, e.playerKilledBy, "minecraft:killed_by", ch, player.Data.Player.Username, "")

				actionTypes := []string{"crafted", "used", "picked_up", "dropped", "broken"}
				for _, actionType := range actionTypes {
					field := fmt.Sprintf("minecraft:%s", actionType)
					e.playerStats(jsonParsed, e.item, field, ch, player.Data.Player.Username, actionType)
				}

				e.playerStatsCustom(jsonParsed, e.animalsBred, "stats.minecraft:custom.minecraft:animals_bred", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.cleanArmor, "stats.minecraft:custom.minecraft:clean_armor", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.cleanBanner, "stats.minecraft:custom.minecraft:clean_banner", ch, player.Data.Player.Username)

				e.playerStatsCustom(jsonParsed, e.openBarrel, "stats.minecraft:custom.minecraft:open_barrel", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.bellRing, "stats.minecraft:custom.minecraft:bell_ring", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.eatCakeSlice, "stats.minecraft:custom.minecraft:eat_cake_slice", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.fillCauldron, "stats.minecraft:custom.minecraft:fill_cauldron", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.openChest, "stats.minecraft:custom.minecraft:open_chest", ch, player.Data.Player.Username)

				damageReceivedTypes := []string{"absorbed", "blocked_by_shield", "resisted", "taken"}
				for _, damageReceivedType := range damageReceivedTypes {
					field := fmt.Sprintf("stats.minecraft:custom.minecraft:damage_%s", damageReceivedType)
					e.playerStatsCustomWithType(jsonParsed, e.damageReceived, field, ch, player.Data.Player.Username, damageReceivedType)
				}

				damageDealtTypes := map[string]string{"dealt": "hit", "dealt_absorbed": "absorbed", "dealt_resisted": "resisted"}
				for key, damageDealtType := range damageDealtTypes {
					field := fmt.Sprintf("stats.minecraft:custom.minecraft:damage_%s", key)
					e.playerStatsCustomWithType(jsonParsed, e.damageDealt, field, ch, player.Data.Player.Username, damageDealtType)
				}

				movementTypes := []string{
					"climb", "crouch", "fall", "fly", "sprint", "swim", "walk", "walk_on_water", "walk_under_water",
					"boat", "aviate", "horse", "minecart", "pig", "strider",
				}
				for _, movementType := range movementTypes {
					field := fmt.Sprintf("stats.minecraft:custom.minecraft:%s_one_cm", movementType)
					e.playerStatsCustomMovement(jsonParsed, e.minecraftMovement, field, ch, player.Data.Player.Username, movementType)
				}

				inspectionTypes := []string{"dispenser", "dropper", "hopper"}
				for _, inspectionType := range inspectionTypes {
					field := fmt.Sprintf("stats.minecraft:custom.minecraft:inspect_%s", inspectionType)
					e.playerStatsCustomWithType(jsonParsed, e.inspected, field, ch, player.Data.Player.Username, inspectionType)
				}

				e.playerStatsCustom(jsonParsed, e.openEnderChest, "stats.minecraft:custom.minecraft:open_enderchest", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.fishCaught, "stats.minecraft:custom.minecraft:fish_caught", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.leaveGame, "stats.minecraft:custom.minecraft:leave_game", ch, player.Data.Player.Username)

				interactionTypes := []string{
					"anvil", "beacon", "blast_furnace", "brewingstand", "campfire", "cartography_table",
					"crafting_table", "furnace", "grindston", "lectern", "loom", "smithing_table", "smoker", "stonecutter",
				}
				for _, interactionType := range interactionTypes {
					field := fmt.Sprintf("stats.minecraft:custom.minecraft:interact_with_%s", interactionType)
					e.playerStatsCustomWithType(jsonParsed, e.interaction, field, ch, player.Data.Player.Username, interactionType)
				}

				e.playerStatsCustom(jsonParsed, e.itemsDropped, "stats.minecraft:custom.minecraft:drop", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.itemsEntchanted, "stats.minecraft:custom.minecraft:enchant_item", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.jump, "stats.minecraft:custom.minecraft:jump", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.mobKills, "stats.minecraft:custom.minecraft:mob_kills", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.musicDiscsPlayed, "stats.minecraft:custom.minecraft:play_record", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.noteBlocksPlayed, "stats.minecraft:custom.minecraft:play_noteblockr", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.noteBlocksTuned, "stats.minecraft:custom.minecraft:tune_noteblock", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.numberOfDeaths, "stats.minecraft:custom.minecraft:deaths", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.plantsPotted, "stats.minecraft:custom.minecraft:pot_flower", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.playerKills, "stats.minecraft:custom.minecraft:player_kills", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.raidsTriggered, "stats.minecraft:custom.minecraft:raid_trigger", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.raidsWon, "stats.minecraft:custom.minecraft:raid_win", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.shulkerBoxCleaned, "stats.minecraft:custom.minecraft:clean_shulker_box", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.shulkerBoxesOpened, "stats.minecraft:custom.minecraft:open_shulker_box", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.sneakTime, "stats.minecraft:custom.minecraft:sneak_time", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.talkedToVillager, "stats.minecraft:custom.minecraft:talked_to_villager", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.targetsHit, "stats.minecraft:custom.minecraft:target_hit", ch, player.Data.Player.Username)

				if pre1_17(jsonParsed, "stats.minecraft:custom.minecraft:play_one_minute") {
					e.playerStatsCustom(jsonParsed, e.timePlayed, "stats.minecraft:custom.minecraft:play_one_minute", ch, player.Data.Player.Username)
				} else {
					e.playerStatsCustom(jsonParsed, e.timePlayed, "stats.minecraft:custom.minecraft:play_time", ch, player.Data.Player.Username)
				}

				e.playerStatsCustom(jsonParsed, e.timeSinceDeath, "stats.minecraft:custom.minecraft:time_since_death", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.timeSinceLastRest, "stats.minecraft:custom.minecraft:time_since_rest", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.timesWorldOpen, "stats.minecraft:custom.minecraft:total_world_time", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.timesSleptInBed, "stats.minecraft:custom.minecraft:sleep_in_bed", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.tradedWithVillagers, "stats.minecraft:custom.minecraft:traded_with_villager", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.trappedChestsTriggered, "stats.minecraft:custom.minecraft:trigger_trapped_chest", ch, player.Data.Player.Username)
				e.playerStatsCustom(jsonParsed, e.waterTakenFromCauldron, "stats.minecraft:custom.minecraft:use_cauldron", ch, player.Data.Player.Username)
			}
		}
	}
	return nil
}

func (e *Exporter) isEnabled(field string) bool {
	for _, group := range strings.Split(field, ".") {
		if value, ok := e.disabledMetrics[group]; ok && value {
			return false
		}
	}

	return true
}

func (e *Exporter) playerStatsCustomWithType(jsonParsed *gabs.Container, desc *prometheus.Desc, field string, ch chan<- prometheus.Metric, playerName, entityType string) {
	if e.isEnabled(field) {
		value, _ := jsonParsed.Path(field).Data().(float64)
		ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, value, playerName, entityType)
	}
}

func (e *Exporter) playerStatsCustomMovement(jsonParsed *gabs.Container, desc *prometheus.Desc, field string, ch chan<- prometheus.Metric, playerName, movementType string) {
	if e.isEnabled(field) {
		value, _ := jsonParsed.Path(field).Data().(float64)
		value /= 100
		ch <- prometheus.MustNewConstMetric(desc, prometheus.UntypedValue, value, playerName, movementType)
	}
}

func pre1_17(jsonParsed *gabs.Container, field string) bool {
	value := jsonParsed.Path(field)
	return value != nil
}

func (e *Exporter) playerStatsCustom(jsonParsed *gabs.Container, desc *prometheus.Desc, field string, ch chan<- prometheus.Metric, playerName string) {
	if e.isEnabled(field) {
		value, _ := jsonParsed.Path(field).Data().(float64)
		ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, value, playerName)
	}
}

func (e *Exporter) playerStats(jsonParsed *gabs.Container, desc *prometheus.Desc, field string, ch chan<- prometheus.Metric, playerName, actionType string) {
	if !e.isEnabled(field) {
		return
	}

	for key, val := range jsonParsed.S("stats", field).ChildrenMap() {
		if !e.isEnabled(key) {
			continue
		}

		val := val.Data().(float64)
		namespace := strings.Split(key, ":")[0]
		entity := strings.Split(key, ":")[1]
		if len(actionType) > 0 {
			ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val, playerName, namespace, entity, actionType)
		} else {
			ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val, playerName, namespace, entity)
		}

	}
}

func (e *Exporter) advancements(id string, ch chan<- prometheus.Metric, playerName string) error {
	var payload map[string]interface{}
	byteValue, err := os.ReadFile(e.world + "/advancements/" + id + ".json")
	if err != nil {
		e.logger.Error(fmt.Sprintf("advancements file for player %s not exist", playerName))
		return nil
	}
	err = json.Unmarshal(byteValue, &payload)
	if err != nil {
		return err
	}
	m := payload
	var completed int
	for _, v := range m {
		if advancement, ok := v.(map[string]interface{}); ok {
			if done, ok := advancement["done"]; ok {
				if done, ok := done.(bool); ok && done {
					completed++
				}
			}
		}
	}
	ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(completed), playerName, "advancements")
	return nil
}

func (e *Exporter) executeRCONCommand(cmd string) (*string, error) {
	if len(e.rcon.rconPassword) > 0 && e.rcon.rconClient == nil {
		e.logger.Warn("RCON is not connected, trying to reconnect")
		e.rcon = createRCONClient(e.rcon.rconServer, e.rcon.rconPassword, e.logger)
	}
	if e.rcon.rconClient != nil {
		err := e.rcon.rconClient.Cmd(cmd)
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				e.logger.Warn("This is broken pipe error, trying to reconnect")
				e.rcon = createRCONClient(e.rcon.rconServer, e.rcon.rconPassword, e.logger)
			}
			return nil, fmt.Errorf("send rcon command error: %w", err)
		}

		resp, err := e.rcon.rconClient.Resp()
		if err != nil {
			return nil, fmt.Errorf("receive rcon command error: %w", err)
		}
		return &resp, nil
	}
	return nil, nil
}

func removeColorCodesFromWord(word string) string {
	if strings.Contains(word, "§") {
		pos := strings.IndexRune(word, '§')
		cleanWord := strings.Split(word, "")
		cleanWord = append(cleanWord[:pos], cleanWord[pos+2:]...)
		return removeColorCodesFromWord(strings.Join(cleanWord, ""))
	}
	return word
}

func (e *Exporter) getPlayerList(ch chan<- prometheus.Metric) (retErr error) {
	resp, err := e.executeRCONCommand(rconListCommand)
	if err != nil {
		return err
	}
	if resp != nil {
		players := e.playerOnlineRegexp.FindStringSubmatch(*resp)
		if len(players) > 1 {
			playersList := strings.TrimSpace(players[1])
			if playersList != "" {
				list := strings.Split(playersList, ",")
				for _, player := range list {
					player = removeColorCodesFromWord(player)
					player = strings.TrimSpace(player)
					if player != "" {
						ch <- prometheus.MustNewConstMetric(e.playerOnline, prometheus.CounterValue, 1, player)
					}
				}
			}
		}
	}
	return nil
}

// just in case there is a different locale
func parseFloat64FromString(value string) float64 {
	if strings.Contains(value, ",") {
		value = strings.ReplaceAll(value, ",", ".")
	}
	valueInFloat64, _ := strconv.ParseFloat(value, 64)
	return valueInFloat64
}

func (e *Exporter) getServerStats(ch chan<- prometheus.Metric) (retErr error) {
	switch e.serverStats {
	case Forge:
		resp, err := e.executeRCONCommand(rconForgeTpsCommand)
		if err != nil {
			return err
		}
		if resp != nil {
			dimTpsList := e.dimensionRegexp.FindAllStringSubmatch(*resp, -1)
			for _, dimTps := range dimTpsList {
				namespace := dimTps[1]
				dimension := dimTps[2]
				meanTickTimeDimension := parseFloat64FromString(dimTps[3])
				meanTpsDimension := parseFloat64FromString(dimTps[4])
				ch <- prometheus.MustNewConstMetric(e.dimTps, prometheus.CounterValue, meanTpsDimension, namespace, dimension)
				ch <- prometheus.MustNewConstMetric(e.dimTicktime, prometheus.CounterValue, meanTickTimeDimension, namespace, dimension)
			}
			overall := e.overallRegexp.FindStringSubmatch(*resp)
			if len(overall) == 3 {
				meanTickTime := parseFloat64FromString(overall[1])
				meanTps := parseFloat64FromString(overall[2])
				ch <- prometheus.MustNewConstMetric(e.overallTps, prometheus.CounterValue, meanTps)
				ch <- prometheus.MustNewConstMetric(e.overallTicktime, prometheus.CounterValue, meanTickTime)
			}
		}
		resp, err = e.executeRCONCommand(rconForgeEntityListCommand)
		if resp != nil {
			if err != nil {
				return err
			}
			entityList := e.entityListRegexp.FindAllStringSubmatch(*resp, -1)
			for _, entity := range entityList {
				entityCounter := parseFloat64FromString(entity[1])
				namespace := entity[2]
				entityType := entity[3]
				ch <- prometheus.MustNewConstMetric(e.entities, prometheus.CounterValue, entityCounter, namespace, entityType)
			}
		}
	case PaperMC, PurpurMC:
		resp, err := e.executeRCONCommand(rconTpsCommand)
		if resp != nil {
			if err != nil {
				return err
			}
			if e.serverStats == PaperMC {
				tpsString := e.paperMcTpsRegexp.FindStringSubmatch(*resp)
				if len(tpsString) == 4 {
					tps := map[float64]uint64{
						1:  uint64(parseFloat64FromString(tpsString[1])),
						5:  uint64(parseFloat64FromString(tpsString[2])),
						15: uint64(parseFloat64FromString(tpsString[3])),
					}
					sum := tps[1] + tps[5] + tps[15]
					ch <- prometheus.MustNewConstHistogram(e.tpsPaperMC, uint64(len(tps)), float64(sum), tps)
				}
			} else {
				tpsString := e.purpurMcTpsRegexp.FindStringSubmatch(*resp)
				if len(tpsString) == 5 {
					tps := map[float64]uint64{
						0.08: uint64(parseFloat64FromString(tpsString[1])),
						1:    uint64(parseFloat64FromString(tpsString[2])),
						5:    uint64(parseFloat64FromString(tpsString[3])),
						15:   uint64(parseFloat64FromString(tpsString[4])),
					}
					sum := tps[0.08] + tps[1] + tps[5] + tps[15]
					ch <- prometheus.MustNewConstHistogram(e.tpsPaperMC, uint64(len(tps)), float64(sum), tps)
				}
			}
		}
	case Fabric:
		resp, err := e.executeRCONCommand(rconFabricTpsCommand)
		if err != nil {
			return err
		}
		if resp != nil {
			dimTpsList := e.dimensionRegexp.FindAllStringSubmatch(*resp, -1)
			for _, dimTps := range dimTpsList {
				namespace := dimTps[1]
				dimension := dimTps[2]
				meanTickTimeDimension := parseFloat64FromString(dimTps[3])
				meanTpsDimension := parseFloat64FromString(dimTps[4])
				ch <- prometheus.MustNewConstMetric(e.dimTps, prometheus.CounterValue, meanTpsDimension, namespace, dimension)
				ch <- prometheus.MustNewConstMetric(e.dimTicktime, prometheus.CounterValue, meanTickTimeDimension, namespace, dimension)
			}
			overall := e.overallRegexp.FindStringSubmatch(*resp)
			if len(overall) == 3 {
				meanTickTime := parseFloat64FromString(overall[1])
				meanTps := parseFloat64FromString(overall[2])
				ch <- prometheus.MustNewConstMetric(e.overallTps, prometheus.CounterValue, meanTps)
				ch <- prometheus.MustNewConstMetric(e.overallTicktime, prometheus.CounterValue, meanTickTime)
			}
		}
	}
	return nil
}

func (e *Exporter) Describe(descs chan<- *prometheus.Desc) {
	descs <- e.playerOnline
	descs <- e.playerStat
	descs <- e.item
	descs <- e.blocksMined
	descs <- e.entitiesKilled
	descs <- e.playerKilledBy
	descs <- e.animalsBred
	descs <- e.cleanArmor
	descs <- e.cleanBanner
	descs <- e.openBarrel
	descs <- e.bellRing
	descs <- e.eatCakeSlice
	descs <- e.fillCauldron
	descs <- e.openChest
	descs <- e.damageDealt
	descs <- e.damageReceived
	descs <- e.inspected
	descs <- e.minecraftMovement
	descs <- e.openEnderChest
	descs <- e.fishCaught
	descs <- e.leaveGame
	descs <- e.interaction
	descs <- e.itemsDropped
	descs <- e.itemsEntchanted
	descs <- e.jump
	descs <- e.mobKills
	descs <- e.musicDiscsPlayed
	descs <- e.noteBlocksPlayed
	descs <- e.noteBlocksTuned
	descs <- e.numberOfDeaths
	descs <- e.plantsPotted
	descs <- e.playerKills
	descs <- e.raidsTriggered
	descs <- e.raidsWon
	descs <- e.shulkerBoxCleaned
	descs <- e.shulkerBoxesOpened
	descs <- e.sneakTime
	descs <- e.talkedToVillager
	descs <- e.targetsHit
	descs <- e.timePlayed
	descs <- e.timeSinceDeath
	descs <- e.timeSinceLastRest
	descs <- e.timesWorldOpen
	descs <- e.timesSleptInBed
	descs <- e.tradedWithVillagers
	descs <- e.trappedChestsTriggered
	descs <- e.waterTakenFromCauldron
}

func (e *Exporter) Collect(metrics chan<- prometheus.Metric) {
	err := e.getPlayerList(metrics)
	if err != nil {
		e.logger.Error("Failed to get player online list", "err", err)
	}

	if err := e.getPlayerStats(metrics); err != nil {
		e.logger.Error("Failed to get player stats", "err", err)
	}

	if err := e.getServerStats(metrics); err != nil {
		e.logger.Error("Failed to get server stats", "err", err)
	}
}

func (e *Exporter) StopRCON() {
	if e.rcon.rconClient != nil {
		e.logger.Info("Stopping RCON")
		err := e.rcon.rconClient.Close()
		if err != nil {
			e.logger.Error("Failed to close RCON connection", "err", err)
		}
	}
}
