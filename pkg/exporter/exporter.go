package exporter

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/Tnze/go-mc/nbt"
	mcnet "github.com/Tnze/go-mc/net"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "minecraft_prometheus_exporter"
)

//See for all details on the statistics of Minecraft https://minecraft.fandom.com/wiki/Statistics

type Exporter struct {
	address  string
	password string
	logger   log.Logger
	world    string
	//via advancements
	playerAdvancements *prometheus.Desc

	//via RCON
	playerOnline *prometheus.Desc

	//via stats
	playerXpTotal   *prometheus.Desc
	playerCurrentXp *prometheus.Desc
	playerFoodLevel *prometheus.Desc
	playerScore     *prometheus.Desc
	playerHealth    *prometheus.Desc
	itemCrafted     *prometheus.Desc
	blocksMined     *prometheus.Desc
	entitiesKilled  *prometheus.Desc
	playerKilledBy  *prometheus.Desc
	itemUsed        *prometheus.Desc
	itemPickedUp    *prometheus.Desc
	itemDropped     *prometheus.Desc
	itemBroken      *prometheus.Desc

	animalsBred                  *prometheus.Desc
	cleanArmor                   *prometheus.Desc
	cleanBanner                  *prometheus.Desc
	openBarrel                   *prometheus.Desc
	bellRing                     *prometheus.Desc
	eatCakeSlice                 *prometheus.Desc
	fillCauldron                 *prometheus.Desc
	openChest                    *prometheus.Desc
	damageAbsorbed               *prometheus.Desc
	damageBlockedByShield        *prometheus.Desc
	damageDealt                  *prometheus.Desc
	damageDealtAbsorbed          *prometheus.Desc
	damageDealtResisted          *prometheus.Desc
	damageResisted               *prometheus.Desc
	damageTaken                  *prometheus.Desc
	inspectDispenser             *prometheus.Desc
	climbOneCm                   *prometheus.Desc
	crouchOneCm                  *prometheus.Desc
	fallOneCm                    *prometheus.Desc
	flyOneCm                     *prometheus.Desc
	sprintOneCm                  *prometheus.Desc
	swimOneCm                    *prometheus.Desc
	walkOneCm                    *prometheus.Desc
	walkOnWaterOneCm             *prometheus.Desc
	walkUnderWaterOneCm          *prometheus.Desc
	boatOneCm                    *prometheus.Desc
	aviateOneCm                  *prometheus.Desc
	horseOneCm                   *prometheus.Desc
	minecartOneCm                *prometheus.Desc
	pigOneCm                     *prometheus.Desc
	striderOneCm                 *prometheus.Desc
	inspectDropper               *prometheus.Desc
	openEnderChest               *prometheus.Desc
	fishCaught                   *prometheus.Desc
	leaveGame                    *prometheus.Desc
	inspectHopper                *prometheus.Desc
	interactWithAnvil            *prometheus.Desc
	interactWithBeacon           *prometheus.Desc
	interactWithBlastFurnace     *prometheus.Desc
	interactWithBrewingStand     *prometheus.Desc
	interactWithCampfire         *prometheus.Desc
	interactWithCartographyTable *prometheus.Desc
	interactWithCraftingTable    *prometheus.Desc
	interactWithFurnaces         *prometheus.Desc
	interactWithGrindstone       *prometheus.Desc
	interactWithLectern          *prometheus.Desc
	interactWithLoom             *prometheus.Desc
	interactWithSmithingTable    *prometheus.Desc
	interactWithSmoker           *prometheus.Desc
	interactWithStonecutter      *prometheus.Desc
	itemsDropped                 *prometheus.Desc
	itemsEntchanted              *prometheus.Desc
	jump                         *prometheus.Desc
	mobKills                     *prometheus.Desc
	musicDiscsPlayed             *prometheus.Desc
	noteBlocksPlayed             *prometheus.Desc
	noteBlocksTuned              *prometheus.Desc
	numberOfDeaths               *prometheus.Desc
	plantsPotted                 *prometheus.Desc
	playerKills                  *prometheus.Desc
	raidsTriggered               *prometheus.Desc
	raidsWon                     *prometheus.Desc
	shulkerBoxCleaned            *prometheus.Desc
	shulkerBoxesOpened           *prometheus.Desc
	sneakTime                    *prometheus.Desc
	talkedToVillager             *prometheus.Desc
	targetsHit                   *prometheus.Desc
	timePlayed                   *prometheus.Desc
	timeSinceDeath               *prometheus.Desc
	timeSinceLastRest            *prometheus.Desc
	timesWorldOpen               *prometheus.Desc
	timesSleptInBed              *prometheus.Desc
	tradedWithVillagers          *prometheus.Desc
	trappedChestsTriggered       *prometheus.Desc
	waterTakenFromCauldron       *prometheus.Desc
}

func New(server, password, world string, logger log.Logger) *Exporter {
	return &Exporter{
		address:  server,
		password: password,
		logger:   logger,
		world:    world,
		playerOnline: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "player_online_total"),
			"is 1 if player is online",
			[]string{"player"},
			nil,
		),
		playerXpTotal: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_xp_total"),
			"How much total XP a player has",
			[]string{"player"},
			nil,
		),
		playerCurrentXp: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_current_xp_total"),
			"How much current XP a player has",
			[]string{"player"},
			nil,
		),
		playerFoodLevel: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_food_level_total"),
			"How much food the player currently has",
			[]string{"player"},
			nil,
		),
		playerHealth: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_health_total"),
			"How much Health the player currently has",
			[]string{"player"},
			nil,
		),
		playerScore: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_score_total"),
			"The Score of the player",
			[]string{"player"},
			nil,
		),
		playerAdvancements: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_advancements_total"),
			"Number of completed advances of a player",
			[]string{"player"},
			nil,
		),
		itemCrafted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "item_crafted_total"),
			"Statistics related to the number of items crafted, smelted, etc.",
			[]string{"player", "block"},
			nil,
		),
		blocksMined: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "blocks_mined_total"),
			"Statistic related to the number of blocks a player mined",
			[]string{"player", "block"},
			nil,
		),
		entitiesKilled: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "entities_killed_total"),
			"Statistics related to the number of entities a player killed",
			[]string{"player", "entity"},
			nil,
		),
		playerKilledBy: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "killed_by_total"),
			"Statistics related to the times of a player being killed by entities.",
			[]string{"player", "entity"},
			nil,
		),
		itemUsed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "item_used_total"),
			"Statistics related to the number of block or item used",
			[]string{"player", "entity"},
			nil,
		),
		itemPickedUp: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "item_picked_up_total"),
			"Statistics related to the number of dropped items a player picked up",
			[]string{"player", "entity"},
			nil,
		),
		itemDropped: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "item_dropped_total"),
			"Statistics related to the number of items that droped.",
			[]string{"player", "entity"},
			nil,
		),
		itemBroken: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "item_broken_total"),
			"Statistics related to the number of items a player ran their durability negative",
			[]string{"player", "entity"},
			nil,
		),
		animalsBred: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "animals_bred_total"),
			"The number of times the player bred two mobs.",
			[]string{"player"},
			nil,
		),
		cleanArmor: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "clean_armor_total"),
			"The number of dyed leather armors washed with a cauldron.",
			[]string{"player"},
			nil,
		),
		//------
		cleanBanner: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "clean_banner_total"),
			"The number of banner patterns washed with a cauldron.",
			[]string{"player"},
			nil,
		),
		openBarrel: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "open_barrel_total"),
			"The number of times the player has opened a Barrel.",
			[]string{"player"},
			nil,
		),
		bellRing: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "bell_ring_total"),
			"The number of times the player has rung a Bell.",
			[]string{"player"},
			nil,
		),
		eatCakeSlice: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "eat_cake_slice_total"),
			"The number of cake slices eaten.",
			[]string{"player"},
			nil,
		),
		fillCauldron: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "fill_cauldron_total"),
			"The number of times the player filled cauldrons with water buckets.",
			[]string{"player"},
			nil,
		),
		openChest: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "open_chest_total"),
			"The number of times the player opened chests.",
			[]string{"player"},
			nil,
		),
		damageAbsorbed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_absorbed_total"),
			"The amount of damage the player has absorbed in tenths of 1♥.",
			[]string{"player"},
			nil,
		),
		damageBlockedByShield: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_blocked_by_shield_total"),
			"The amount of damage the player has blocked with a shield in tenths of 1♥.",
			[]string{"player"},
			nil,
		),
		damageDealt: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_dealt_total"),
			"The amount of damage the player has dealt in tenths 1♥. Includes only melee attacks.",
			[]string{"player"},
			nil,
		),
		damageDealtAbsorbed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_dealt_absorbed_total"),
			"The amount of damage the player has dealt that were absorbed, in tenths of 1♥.",
			[]string{"player"},
			nil,
		),
		damageDealtResisted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_dealt_resisted_total"),
			"The amount of damage the player has dealt that were resisted, in tenths of 1♥.",
			[]string{"player"},
			nil,
		),
		damageResisted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_resisted_total"),
			"The amount of damage the player has resisted in tenths of 1♥.",
			[]string{"player"},
			nil,
		),
		damageTaken: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "damage_taken_total"),
			"The amount of damage the player has taken in tenths of 1♥.",
			[]string{"player"},
			nil,
		),
		inspectDispenser: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "inspect_dispenser_total"),
			"The number of times interacted with dispensers.",
			[]string{"player"},
			nil,
		),
		climbOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "climb_one_cm_total"),
			"The total distance traveled up ladders or vines.",
			[]string{"player"},
			nil,
		),
		crouchOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "crouch_one_cm_total"),
			"The total distance walked while sneaking.",
			[]string{"player"},
			nil,
		),
		fallOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "fall_one_cm_total"),
			"The total distance fallen, excluding jumping. ",
			[]string{"player"},
			nil,
		),
		flyOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "fly_one_cm_total"),
			"Distance traveled upwards and forwards at the same time, while more than one block above the ground.",
			[]string{"player"},
			nil,
		),
		sprintOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "sprint_one_cm_total"),
			"The total distance sprinted.",
			[]string{"player"},
			nil,
		),
		swimOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "swim_one_cm_total"),
			"The total distance covered with sprint-swimming..",
			[]string{"player"},
			nil,
		),
		walkOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "walk_one_cm_total"),
			"The total distance walked.",
			[]string{"player"},
			nil,
		),
		walkOnWaterOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "walk_on_water_one_cm_total"),
			"The distance covered while bobbing up and down over water.",
			[]string{"player"},
			nil,
		),
		walkUnderWaterOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "walk_under_water_one_cm_total"),
			"The total distance you have walked underwater.",
			[]string{"player"},
			nil,
		),
		boatOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "boat_one_cm_total"),
			"The total distance traveled by boats.",
			[]string{"player"},
			nil,
		),
		aviateOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "aviate_one_cm_total"),
			"The total distance traveled by elytra.",
			[]string{"player"},
			nil,
		),
		horseOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "horse_one_cm_total"),
			"The total distance traveled by horses..",
			[]string{"player"},
			nil,
		),
		minecartOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "minecart_one_cm_total"),
			"The total distance traveled by minecarts.",
			[]string{"player"},
			nil,
		),
		pigOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "pig_one_cm_total"),
			"The total distance traveled by pigs via saddles.",
			[]string{"player"},
			nil,
		),
		striderOneCm: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "strider_one_cm_total"),
			"The total distance traveled by striders via saddles.",
			[]string{"player"},
			nil,
		),
		inspectDropper: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "inspect_dropper_total"),
			"The number of times interacted with droppers.",
			[]string{"player"},
			nil,
		),
		openEnderChest: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "open_enderchest_total"),
			"The number of times the player opened ender chests.",
			[]string{"player"},
			nil,
		),
		fishCaught: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "fish_caught_total"),
			"The number of fish caught.",
			[]string{"player"},
			nil,
		),
		leaveGame: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "leave_game_total"),
			"The number of times \"Save and quit to title\" has been clicked.",
			[]string{"player"},
			nil,
		),
		inspectHopper: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "inspect_hopper_total"),
			"The number of times interacted with hoppers.",
			[]string{"player"},
			nil,
		),
		interactWithAnvil: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_anvil_total"),
			"The number of times interacted with anvils.",
			[]string{"player"},
			nil,
		),
		interactWithBeacon: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_beacon_total"),
			"The number of times interacted with beacons.",
			[]string{"player"},
			nil,
		),
		interactWithBlastFurnace: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_blast_furnace_total"),
			"The number of times interacted with blast furnaces",
			[]string{"player"},
			nil,
		),
		interactWithBrewingStand: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_brewingstand_total"),
			"The number of times interacted with brewing stands",
			[]string{"player"},
			nil,
		),
		interactWithCampfire: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_campfire_total"),
			"The number of times interacted with campfires",
			[]string{"player"},
			nil,
		),
		interactWithCartographyTable: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_cartography_table_total"),
			"The number of times interacted with cartography tables",
			[]string{"player"},
			nil,
		),
		interactWithCraftingTable: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_crafting_table_total"),
			"The number of times interacted with crafting tables",
			[]string{"player"},
			nil,
		),
		interactWithFurnaces: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_furnace_total"),
			"The number of times interacted with furnaces",
			[]string{"player"},
			nil,
		),
		interactWithGrindstone: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_grindstone_total"),
			"The number of times interacted with grindstones",
			[]string{"player"},
			nil,
		),
		interactWithLectern: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_lectern_total"),
			"The number of times interacted with lecterns",
			[]string{"player"},
			nil,
		),
		interactWithLoom: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_loom_total"),
			"The number of times interacted with looms",
			[]string{"player"},
			nil,
		),
		interactWithSmithingTable: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_smithing_table_total"),
			"The number of times interacted with smithing tables.",
			[]string{"player"},
			nil,
		),
		interactWithSmoker: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_smoker_total"),
			"The number of times interacted with smokers.",
			[]string{"player"},
			nil,
		),
		interactWithStonecutter: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "interact_with_stonecutter_total"),
			"The number of times interacted with stonecutters.",
			[]string{"player"},
			nil,
		),
		itemsDropped: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "items_drop_total"),
			"The number of items dropped.",
			[]string{"player"},
			nil,
		),
		itemsEntchanted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "enchant_item_total"),
			"The number of items enchanted.",
			[]string{"player"},
			nil,
		),
		jump: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "jump_total"),
			"\tThe total number of jumps performed.",
			[]string{"player"},
			nil,
		),
		mobKills: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "mob_kills_total"),
			"The number of mobs the player killed.",
			[]string{"player"},
			nil,
		),
		musicDiscsPlayed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "play_record_total"),
			"The number of music discs played on a jukebox.",
			[]string{"player"},
			nil,
		),
		noteBlocksPlayed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "play_noteblock_total"),
			"The number of note blocks hit.",
			[]string{"player"},
			nil,
		),
		noteBlocksTuned: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "tune_noteblock_total"),
			"The number of times interacted with note blocks.",
			[]string{"player"},
			nil,
		),
		numberOfDeaths: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "deaths_total"),
			"The number of times the player died.",
			[]string{"player"},
			nil,
		),
		plantsPotted: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "pot_flower_total"),
			"The number of plants potted onto flower pots.",
			[]string{"player"},
			nil,
		),
		playerKills: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "player_kills_total"),
			"The number of players the player killed",
			[]string{"player"},
			nil,
		),
		raidsTriggered: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "raid_trigger_total"),
			"The number of times the player has triggered a Raid.",
			[]string{"player"},
			nil,
		),
		raidsWon: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "raid_win_total"),
			"The number of times the player has won a Raid.",
			[]string{"player"},
			nil,
		),
		shulkerBoxCleaned: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "clean_shulker_box_total"),
			"The number of times the player has washed a Shulker Box with a cauldron.",
			[]string{"player"},
			nil,
		),
		shulkerBoxesOpened: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "open_shulker_box_total"),
			"The number of times the player has opened a Shulker Box.",
			[]string{"player"},
			nil,
		),
		sneakTime: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "sneak_time_total"),
			"The time the player has held down the sneak button.",
			[]string{"player"},
			nil,
		),
		talkedToVillager: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "talked_to_villager_total"),
			"The number of times interacted with villagers (opened the trading GUI).",
			[]string{"player"},
			nil,
		),
		targetsHit: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "target_hit_total"),
			"The number of times the player has shot a target block.",
			[]string{"player"},
			nil,
		),
		timePlayed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "play_time_total"),
			"The total amount of time played. ",
			[]string{"player"},
			nil,
		),
		timeSinceDeath: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "time_since_death_total"),
			"The time since the player's last death.",
			[]string{"player"},
			nil,
		),
		timeSinceLastRest: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "time_since_rest_total"),
			"The time since the player's last rest. This is used to spawn phantoms.",
			[]string{"player"},
			nil,
		),
		timesWorldOpen: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "total_world_time_total"),
			"The total amount of time the world was opened.n.",
			[]string{"player"},
			nil,
		),
		timesSleptInBed: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "sleep_in_bed_total"),
			"The number of times the player has slept in a bed..",
			[]string{"player"},
			nil,
		),
		tradedWithVillagers: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "traded_with_villager_total"),
			"The number of times traded with villagers.",
			[]string{"player"},
			nil,
		),
		trappedChestsTriggered: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "trigger_trapped_chest_total"),
			"The number of times the player opened trapped chests.",
			[]string{"player"},
			nil,
		),
		waterTakenFromCauldron: prometheus.NewDesc(prometheus.BuildFQName(Namespace, "", "use_cauldron_total"),
			"The number of times the player took water from cauldrons with glass bottles.",
			[]string{"player"},
			nil,
		),
	}
}

type Player struct {
	ID   string
	Name string `json:"name"`
}

type PlayerData struct {
	XpLevel   int32
	XpTotal   int32
	Score     int32
	Health    float32
	FoodLevel int32 `nbt:"foodLevel"`
}

func (e *Exporter) getPlayerStats(ch chan<- prometheus.Metric) error {
	files, err := ioutil.ReadDir(e.world + "/stats")
	if err != nil {
		return err
	}
	for _, file := range files {
		id := strings.TrimSuffix(file.Name(), ".json")
		URL := "https://api.mojang.com/user/profiles/" + id + "/names"
		resp, err := http.Get(URL)
		if err != nil {
			level.Error(e.logger).Log("msg", "Failed to connect to api.mojang.com", "err", err)
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("error retrieving player info from api.mojang.com: %w", errors.New(fmt.Sprintf("Status Code: %d", resp.StatusCode)))
		}

		var cResp []Player

		if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
			level.Error(e.logger).Log("msg", "Failed to connect decode response", "err", err)
		}
		cResp[0].ID = id
		cResp[0].Name = strings.TrimSpace(cResp[0].Name)

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

		ch <- prometheus.MustNewConstMetric(e.playerXpTotal, prometheus.CounterValue, float64(data.XpTotal), cResp[0].Name)
		ch <- prometheus.MustNewConstMetric(e.playerCurrentXp, prometheus.CounterValue, float64(data.XpLevel), cResp[0].Name)
		ch <- prometheus.MustNewConstMetric(e.playerScore, prometheus.CounterValue, float64(data.Score), cResp[0].Name)
		ch <- prometheus.MustNewConstMetric(e.playerFoodLevel, prometheus.CounterValue, float64(data.FoodLevel), cResp[0].Name)
		ch <- prometheus.MustNewConstMetric(e.playerHealth, prometheus.CounterValue, float64(data.Health), cResp[0].Name)

		err = resp.Body.Close()
		if err != nil {
			return err
		}

		err2 := e.advancements(id, ch, cResp[0].Name)
		if err2 != nil {
			return err2
		}

		stats, err := os.Open(e.world + "/stats/" + id + ".json")
		if err != nil {
			return err
		}

		byteValue, _ := ioutil.ReadAll(stats)
		jsonParsed, err := gabs.ParseJSON(byteValue)
		if err != nil {
			return err
		}

		e.playerStats(jsonParsed, e.itemCrafted, "minecraft:crafted", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.blocksMined, "minecraft:mined", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.entitiesKilled, "minecraft:killed", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.playerKilledBy, "minecraft:killed_by", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.itemUsed, "minecraft:used", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.itemPickedUp, "minecraft:picked_up", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.itemDropped, "minecraft:dropped", ch, cResp[0].Name)
		e.playerStats(jsonParsed, e.itemBroken, "minecraft:broken", ch, cResp[0].Name)

		e.playerStatsCustom(jsonParsed, e.animalsBred, "stats.minecraft:custom.minecraft:animals_bred", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.cleanArmor, "stats.minecraft:custom.minecraft:clean_armor", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.cleanBanner, "stats.minecraft:custom.minecraft:clean_armor", ch, cResp[0].Name)

		e.playerStatsCustom(jsonParsed, e.openBarrel, "stats.minecraft:custom.minecraft:open_barrel", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.bellRing, "stats.minecraft:custom.minecraft:bell_ring", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.eatCakeSlice, "stats.minecraft:custom.minecraft:eat_cake_slice", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.fillCauldron, "stats.minecraft:custom.minecraft:fill_cauldron", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.openChest, "stats.minecraft:custom.minecraft:open_chest", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageAbsorbed, "stats.minecraft:custom.minecraft:damage_absorbed", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageBlockedByShield, "stats.minecraft:custom.minecraft:damage_blocked_by_shield", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageDealt, "stats.minecraft:custom.minecraft:damage_dealt", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageDealtAbsorbed, "stats.minecraft:custom.minecraft:damage_dealt_absorbed", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageDealtResisted, "stats.minecraft:custom.minecraft:damage_dealt_resisted", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageResisted, "stats.minecraft:custom.minecraft:damage_resisted", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.damageTaken, "stats.minecraft:custom.minecraft:damage_taken", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.inspectDispenser, "stats.minecraft:custom.minecraft:inspect_dispenserr", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.climbOneCm, "stats.minecraft:custom.minecraft:climb_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.crouchOneCm, "stats.minecraft:custom.minecraft:crouch_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.fallOneCm, "stats.minecraft:custom.minecraft:fall_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.flyOneCm, "stats.minecraft:custom.minecraft:fly_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.sprintOneCm, "stats.minecraft:custom.minecraft:sprint_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.swimOneCm, "stats.minecraft:custom.minecraft:swim_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.walkOneCm, "stats.minecraft:custom.minecraft:walk_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.walkOnWaterOneCm, "stats.minecraft:custom.minecraft:walk_on_water_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.walkUnderWaterOneCm, "stats.minecraft:custom.minecraft:walk_under_water_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.boatOneCm, "stats.minecraft:custom.minecraft:boat_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.aviateOneCm, "stats.minecraft:custom.minecraft:aviate_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.horseOneCm, "stats.minecraft:custom.minecraft:horse_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.minecartOneCm, "stats.minecraft:custom.minecraft:minecart_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.pigOneCm, "stats.minecraft:custom.minecraft:pig_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.striderOneCm, "stats.minecraft:custom.minecraft:strider_one_cm", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.inspectDropper, "stats.minecraft:custom.minecraft:inspect_dropper", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.openEnderChest, "stats.minecraft:custom.minecraft:open_enderchest", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.fishCaught, "stats.minecraft:custom.minecraft:fish_caught", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.leaveGame, "stats.minecraft:custom.minecraft:leave_game", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.inspectHopper, "stats.minecraft:custom.minecraft:inspect_hopper", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithAnvil, "stats.minecraft:custom.minecraft:interact_with_anvil", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithBeacon, "stats.minecraft:custom.minecraft:interact_with_beacon", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithBlastFurnace, "stats.minecraft:custom.minecraft:interact_with_blast_furnace", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithBrewingStand, "stats.minecraft:custom.minecraft:interact_with_brewingstand", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithCampfire, "stats.minecraft:custom.minecraft:interact_with_campfire", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithCartographyTable, "stats.minecraft:custom.minecraft:interact_with_cartography_table", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithCraftingTable, "stats.minecraft:custom.minecraft:interact_with_crafting_table", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithFurnaces, "stats.minecraft:custom.minecraft:interact_with_furnace", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithGrindstone, "stats.minecraft:custom.minecraft:interact_with_grindstone", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithLectern, "stats.minecraft:custom.minecraft:interact_with_lectern", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithLoom, "stats.minecraft:custom.minecraft:interact_with_loom", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithSmithingTable, "stats.minecraft:custom.minecraft:interact_with_smithing_table", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithSmoker, "stats.minecraft:custom.minecraft:interact_with_smoker", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.interactWithStonecutter, "stats.minecraft:custom.minecraft:interact_with_stonecutter", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.itemsDropped, "stats.minecraft:custom.minecraft:drop", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.itemsEntchanted, "stats.minecraft:custom.minecraft:enchant_item", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.jump, "stats.minecraft:custom.minecraft:jump", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.mobKills, "stats.minecraft:custom.minecraft:mob_kills", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.musicDiscsPlayed, "stats.minecraft:custom.minecraft:play_record", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.noteBlocksPlayed, "stats.minecraft:custom.minecraft:play_noteblockr", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.noteBlocksTuned, "stats.minecraft:custom.minecraft:tune_noteblock", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.numberOfDeaths, "stats.minecraft:custom.minecraft:deaths", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.plantsPotted, "stats.minecraft:custom.minecraft:pot_flower", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.playerKills, "stats.minecraft:custom.minecraft:player_kills", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.raidsTriggered, "stats.minecraft:custom.minecraft:raid_trigger", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.raidsWon, "stats.minecraft:custom.minecraft:raid_win", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.shulkerBoxCleaned, "stats.minecraft:custom.minecraft:clean_shulker_box", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.shulkerBoxesOpened, "stats.minecraft:custom.minecraft:open_shulker_box", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.sneakTime, "stats.minecraft:custom.minecraft:sneak_time", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.talkedToVillager, "stats.minecraft:custom.minecraft:talked_to_villager", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.targetsHit, "stats.minecraft:custom.minecraft:target_hit", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.timePlayed, "stats.minecraft:custom.minecraft:play_timer", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.timeSinceDeath, "stats.minecraft:custom.minecraft:time_since_death", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.timeSinceLastRest, "stats.minecraft:custom.minecraft:time_since_rest", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.timesWorldOpen, "stats.minecraft:custom.minecraft:total_world_time", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.timesSleptInBed, "stats.minecraft:custom.minecraft:sleep_in_bed", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.tradedWithVillagers, "stats.minecraft:custom.minecraft:traded_with_villager", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.trappedChestsTriggered, "stats.minecraft:custom.minecraft:trigger_trapped_chest", ch, cResp[0].Name)
		e.playerStatsCustom(jsonParsed, e.waterTakenFromCauldron, "stats.minecraft:custom.minecraft:use_cauldron", ch, cResp[0].Name)
	}
	return nil
}

func (e *Exporter) playerStatsCustom(jsonParsed *gabs.Container, desc *prometheus.Desc, field string, ch chan<- prometheus.Metric, playerName string) {
	value, _ := jsonParsed.Path(field).Data().(float64)
	ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, value, playerName)
}

func (e *Exporter) playerStats(jsonParsed *gabs.Container, desc *prometheus.Desc, field string, ch chan<- prometheus.Metric, playerName string) {
	for key, val := range jsonParsed.S("stats", field).ChildrenMap() {
		val := val.Data().(float64)
		entity := strings.Split(key, ":")[1]
		ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, val, playerName, entity)
	}
}

func (e *Exporter) advancements(id string, ch chan<- prometheus.Metric, playerName string) error {
	advancements, err := os.Open(e.world + "/advancements/" + id + ".json")
	if err != nil {
		return err
	}
	var payload map[string]interface{}
	byteValue, err := ioutil.ReadAll(advancements)
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
	ch <- prometheus.MustNewConstMetric(e.playerAdvancements, prometheus.CounterValue, float64(completed), playerName)

	return nil
}

func (e *Exporter) Describe(descs chan<- *prometheus.Desc) {
	descs <- e.playerAdvancements
	descs <- e.playerOnline
	descs <- e.playerXpTotal
	descs <- e.playerCurrentXp
	descs <- e.playerFoodLevel
	descs <- e.playerScore
	descs <- e.playerHealth
	descs <- e.itemCrafted
	descs <- e.blocksMined
	descs <- e.entitiesKilled
	descs <- e.playerKilledBy
	descs <- e.itemUsed
	descs <- e.itemPickedUp
	descs <- e.itemDropped
	descs <- e.itemBroken

	descs <- e.animalsBred
	descs <- e.cleanArmor
	descs <- e.cleanBanner
	descs <- e.openBarrel
	descs <- e.bellRing
	descs <- e.eatCakeSlice
	descs <- e.fillCauldron
	descs <- e.openChest
	descs <- e.damageAbsorbed
	descs <- e.damageBlockedByShield
	descs <- e.damageDealt
	descs <- e.damageDealtAbsorbed
	descs <- e.damageDealtResisted
	descs <- e.damageResisted
	descs <- e.damageTaken
	descs <- e.inspectDispenser
	descs <- e.climbOneCm
	descs <- e.crouchOneCm
	descs <- e.fallOneCm
	descs <- e.flyOneCm
	descs <- e.sprintOneCm
	descs <- e.swimOneCm
	descs <- e.walkOneCm
	descs <- e.walkOnWaterOneCm
	descs <- e.walkUnderWaterOneCm
	descs <- e.boatOneCm
	descs <- e.aviateOneCm
	descs <- e.horseOneCm
	descs <- e.minecartOneCm
	descs <- e.pigOneCm
	descs <- e.striderOneCm
	descs <- e.inspectDropper
	descs <- e.openEnderChest
	descs <- e.fishCaught
	descs <- e.leaveGame
	descs <- e.inspectHopper
	descs <- e.interactWithAnvil
	descs <- e.interactWithBeacon
	descs <- e.interactWithBlastFurnace
	descs <- e.interactWithBrewingStand
	descs <- e.interactWithCampfire
	descs <- e.interactWithCartographyTable
	descs <- e.interactWithCraftingTable
	descs <- e.interactWithFurnaces
	descs <- e.interactWithGrindstone
	descs <- e.interactWithLectern
	descs <- e.interactWithLoom
	descs <- e.interactWithSmithingTable
	descs <- e.interactWithSmoker
	descs <- e.interactWithStonecutter
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

func (e *Exporter) getPlayerList(ch chan<- prometheus.Metric) (retErr error) {
	conn, err := mcnet.DialRCON(e.address, e.password)
	if err != nil {
		return fmt.Errorf("connect rcon error: %w", err)

	}

	defer func() {
		err := conn.Close()
		if err != nil {
			level.Error(e.logger).Log("msg", "Failed to close rcon endpoint", "err", err)
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

func (e *Exporter) Collect(metrics chan<- prometheus.Metric) {
	if len(e.password) > 0 {
		err := e.getPlayerList(metrics)
		if err != nil {
			metrics <- prometheus.MustNewConstMetric(e.playerOnline, prometheus.CounterValue, 0, "")
			level.Error(e.logger).Log("msg", "Failed to get player online list", "err", err)
		}
	}

	if err := e.getPlayerStats(metrics); err != nil {
		level.Error(e.logger).Log("msg", "Failed to get player stats", "err", err)
	}
}
