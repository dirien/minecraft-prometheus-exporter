package exporter

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-kit/log"

	"github.com/Tnze/go-mc/nbt"
	mcnet "github.com/Tnze/go-mc/net"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "minecraft"
)

// See for all details on the statistics of Minecraft https://minecraft.fandom.com/wiki/Statistics

type Exporter struct {
	address         string
	password        string
	logger          log.Logger
	world           string
	source          string
	disabledMetrics map[string]bool
	// via advancements
	// playerAdvancements *prometheus.Desc

	// via RCON
	playerOnline *prometheus.Desc

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

func New(server, password, world, source string, disabledMetrics map[string]bool, logger log.Logger) *Exporter {
	return &Exporter{
		address:         server,
		password:        password,
		logger:          logger,
		world:           world,
		source:          source,
		disabledMetrics: disabledMetrics,
		playerOnline: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "player_online_total"),
			"is 1 if player is online",
			[]string{"player"},
			nil,
		),
		playerStat: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_stat_total"),
			"Different stats of the player: xp, current_xp, food_level, health, score, advancements",
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
			"The total distance the player traveled with different entities (ladders, boats, etc.)",
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
			"The number of times the player caught",
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
			"The number of ticks the world has been loaded",
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
	}
}

type Player struct {
	ID   string `json:"uuid"`
	Name string `json:"username"`
}

type PlayerData struct {
	XpLevel   int32
	XpTotal   int32
	Score     int32
	Health    float32
	FoodLevel int32 `nbt:"foodLevel"`
	Bukkit    struct {
		LastKnownName string `nbt:"lastKnownName"`
	} `nbt:"bukkit"`
}

func (e *Exporter) getPlayerStats(ch chan<- prometheus.Metric) error {
	files, err := os.ReadDir(e.world + "/stats")
	if err != nil {
		return err
	}
	for _, file := range files {
		id := strings.TrimSuffix(file.Name(), ".json")

		f, err := os.Open(e.world + "/playerdata/" + id + ".dat")
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

		var player Player
		switch e.source {
		case "mojang":
			resp, err := http.Get(fmt.Sprintf("https://api.ashcon.app/mojang/v2/user/%s", id))
			if err != nil {
				level.Error(e.logger).Log("msg", "Failed to connect to api.ashcon.app", "err", err) // nolint: errcheck
			}

			if resp.StatusCode == 200 {
				if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
					level.Error(e.logger).Log("msg", "Failed to connect decode response", "err", err) // nolint: errcheck
				}
			} else {
				return fmt.Errorf("error retrieving player info from api.ashcon.app: %w", fmt.Errorf(fmt.Sprintf("Status Code: %d", resp.StatusCode)))
			}

			err = resp.Body.Close()
			if err != nil {
				return err
			}
		case "bukkit":
			if data.Bukkit.LastKnownName == "" {
				return fmt.Errorf("error retrieving player info from nbt: bukkit name is unknown")
			}

			player = Player{
				ID:   id,
				Name: data.Bukkit.LastKnownName,
			}
		default:
			player = Player{
				ID:   id,
				Name: id,
			}
		}

		ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.XpTotal), player.Name, "xp")
		ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.XpLevel), player.Name, "current_xp")
		ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.Score), player.Name, "score")
		ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.FoodLevel), player.Name, "food_level")
		ch <- prometheus.MustNewConstMetric(e.playerStat, prometheus.GaugeValue, float64(data.Health), player.Name, "health")

		err = e.advancements(id, ch, player.Name)
		if err != nil {
			return err
		}

		byteValue, err := os.ReadFile(e.world + "/stats/" + id + ".json")
		if err != nil {
			return err
		}
		jsonParsed, err := gabs.ParseJSON(byteValue)
		if err != nil {
			return err
		}

		e.playerStats(jsonParsed, e.blocksMined, "minecraft:mined", ch, player.Name, "")
		e.playerStats(jsonParsed, e.entitiesKilled, "minecraft:killed", ch, player.Name, "")
		e.playerStats(jsonParsed, e.playerKilledBy, "minecraft:killed_by", ch, player.Name, "")

		actionTypes := []string{"crafted", "used", "picked_up", "dropped", "broken"}
		for _, actionType := range actionTypes {
			field := fmt.Sprintf("minecraft:%s", actionType)
			e.playerStats(jsonParsed, e.item, field, ch, player.Name, actionType)
		}

		e.playerStatsCustom(jsonParsed, e.animalsBred, "stats.minecraft:custom.minecraft:animals_bred", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.cleanArmor, "stats.minecraft:custom.minecraft:clean_armor", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.cleanBanner, "stats.minecraft:custom.minecraft:clean_armor", ch, player.Name)

		e.playerStatsCustom(jsonParsed, e.openBarrel, "stats.minecraft:custom.minecraft:open_barrel", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.bellRing, "stats.minecraft:custom.minecraft:bell_ring", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.eatCakeSlice, "stats.minecraft:custom.minecraft:eat_cake_slice", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.fillCauldron, "stats.minecraft:custom.minecraft:fill_cauldron", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.openChest, "stats.minecraft:custom.minecraft:open_chest", ch, player.Name)

		damageReceivedTypes := []string{"absorbed", "blocked_by_shield", "resisted", "taken"}
		for _, damageReceivedType := range damageReceivedTypes {
			field := fmt.Sprintf("stats.minecraft:custom.minecraft:damage_%s", damageReceivedType)
			e.playerStatsCustomWithType(jsonParsed, e.damageReceived, field, ch, player.Name, damageReceivedType)
		}

		damageDealtTypes := map[string]string{"dealt": "hit", "dealt_absorbed": "absorbed", "dealt_resisted": "resisted"}
		for key, damageDealtType := range damageDealtTypes {
			field := fmt.Sprintf("stats.minecraft:custom.minecraft:damage_%s", key)
			e.playerStatsCustomWithType(jsonParsed, e.damageDealt, field, ch, player.Name, damageDealtType)
		}

		movementTypes := []string{
			"climb", "crouch", "fall", "fly", "sprint", "swim", "walk", "walk_on_water", "walk_under_water",
			"boat", "aviate", "horse", "minecart", "pig", "strider",
		}
		for _, movementType := range movementTypes {
			field := fmt.Sprintf("stats.minecraft:custom.minecraft:%s_one_cm", movementType)
			e.playerStatsCustomMovement(jsonParsed, e.minecraftMovement, field, ch, player.Name, movementType)
		}

		inspectionTypes := []string{"dispenser", "dropper", "hopper"}
		for _, inspectionType := range inspectionTypes {
			field := fmt.Sprintf("stats.minecraft:custom.minecraft:inspect_%s", inspectionType)
			e.playerStatsCustomWithType(jsonParsed, e.inspected, field, ch, player.Name, inspectionType)
		}

		e.playerStatsCustom(jsonParsed, e.openEnderChest, "stats.minecraft:custom.minecraft:open_enderchest", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.fishCaught, "stats.minecraft:custom.minecraft:fish_caught", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.leaveGame, "stats.minecraft:custom.minecraft:leave_game", ch, player.Name)

		interactionTypes := []string{
			"anvil", "beacon", "blast_furnace", "brewingstand", "campfire", "cartography_table",
			"crafting_table", "furnace", "grindston", "lectern", "loom", "smithing_table", "smoker", "stonecutter",
		}
		for _, interactionType := range interactionTypes {
			field := fmt.Sprintf("stats.minecraft:custom.minecraft:interact_with_%s", interactionType)
			e.playerStatsCustomWithType(jsonParsed, e.interaction, field, ch, player.Name, interactionType)
		}

		e.playerStatsCustom(jsonParsed, e.itemsDropped, "stats.minecraft:custom.minecraft:drop", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.itemsEntchanted, "stats.minecraft:custom.minecraft:enchant_item", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.jump, "stats.minecraft:custom.minecraft:jump", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.mobKills, "stats.minecraft:custom.minecraft:mob_kills", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.musicDiscsPlayed, "stats.minecraft:custom.minecraft:play_record", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.noteBlocksPlayed, "stats.minecraft:custom.minecraft:play_noteblockr", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.noteBlocksTuned, "stats.minecraft:custom.minecraft:tune_noteblock", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.numberOfDeaths, "stats.minecraft:custom.minecraft:deaths", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.plantsPotted, "stats.minecraft:custom.minecraft:pot_flower", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.playerKills, "stats.minecraft:custom.minecraft:player_kills", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.raidsTriggered, "stats.minecraft:custom.minecraft:raid_trigger", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.raidsWon, "stats.minecraft:custom.minecraft:raid_win", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.shulkerBoxCleaned, "stats.minecraft:custom.minecraft:clean_shulker_box", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.shulkerBoxesOpened, "stats.minecraft:custom.minecraft:open_shulker_box", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.sneakTime, "stats.minecraft:custom.minecraft:sneak_time", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.talkedToVillager, "stats.minecraft:custom.minecraft:talked_to_villager", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.targetsHit, "stats.minecraft:custom.minecraft:target_hit", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.timePlayed, "stats.minecraft:custom.minecraft:play_timer", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.timeSinceDeath, "stats.minecraft:custom.minecraft:time_since_death", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.timeSinceLastRest, "stats.minecraft:custom.minecraft:time_since_rest", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.timesWorldOpen, "stats.minecraft:custom.minecraft:total_world_time", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.timesSleptInBed, "stats.minecraft:custom.minecraft:sleep_in_bed", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.tradedWithVillagers, "stats.minecraft:custom.minecraft:traded_with_villager", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.trappedChestsTriggered, "stats.minecraft:custom.minecraft:trigger_trapped_chest", ch, player.Name)
		e.playerStatsCustom(jsonParsed, e.waterTakenFromCauldron, "stats.minecraft:custom.minecraft:use_cauldron", ch, player.Name)
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
		return err
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

func (e *Exporter) getPlayerList(ch chan<- prometheus.Metric) (retErr error) {
	conn, err := mcnet.DialRCON(e.address, e.password)
	if err != nil {
		return fmt.Errorf("connect rcon error: %w", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			level.Error(e.logger).Log("msg", "Failed to close rcon endpoint", "err", err) // nolint: errcheck
			if retErr == nil {
				retErr = err
			}
		}
	}()

	err = conn.Cmd("list")
	if err != nil {
		return fmt.Errorf("send rcon command error: %w", err)
	}

	resp, err := conn.Resp()
	if err != nil {
		return fmt.Errorf("receive rcon command error: %w", err)
	}

	r := regexp.MustCompile("players online:(.*)")
	playersraw := r.FindStringSubmatch(resp)[1]
	for _, player := range strings.Fields(strings.ReplaceAll(playersraw, ",", " ")) {
		ch <- prometheus.MustNewConstMetric(e.playerOnline, prometheus.CounterValue, 1, strings.TrimSpace(player))
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
	if len(e.password) > 0 {
		err := e.getPlayerList(metrics)
		if err != nil {
			metrics <- prometheus.MustNewConstMetric(e.playerOnline, prometheus.CounterValue, 0, "")
			level.Error(e.logger).Log("msg", "Failed to get player online list", "err", err) // nolint: errcheck
		}
	}

	if err := e.getPlayerStats(metrics); err != nil {
		level.Error(e.logger).Log("msg", "Failed to get player stats", "err", err) // nolint: errcheck
	}
}
