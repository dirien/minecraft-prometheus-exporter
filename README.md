###### ⚠️ Statistics are present only in Java Edition. Bedrock Edition has no equivalent of statistics in-game ⚠️

# Minecraft Exporter for Prometheus

![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)
![Minecraft](https://img.shields.io/badge/Minecraft-62B47A?style=for-the-badge&logo=Minecraft&logoColor=white)
![Grafana](https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=Grafana&logoColor=white)

This is a Prometheus Minecraft exporter, created as part of the [minectl 🗺](https://github.com/dirien/minectl) project.

It collects metrics from different sources of the game

- RCON
- NBT Files
- Advancement file
- Player stats file

#### RCON

RCON is a protocol that can be used to remotely execute commands to your Minecraft server.

To enable rcon on your minecraft server add the following to the `java` field in your `minectl` manifest:

```yaml
    ...
    java:
      xmx: 2G
      xms: 2G
      rcon:
        password: test
        port: 25575
        enabled: true
        broadcast: true
    edition: java
    ...
```

See [server config examples](https://github.com/dirien/minectl#server-config-) for details and usage of the minectl 🗺
manifest file.

Alternatively you need to edit your server.properties on your server.

Be sure that following flags are set for the `minecraft-exporter` binary:

```bash
--mc.rcon-address=":25575"  
--mc.rcon-password=<rcon-password>
```

#### Files

To access the files, take care that the `minecraft-exporter` binary is started with the flag

Details for the specific stats can be found here -> https://minecraft.fandom.com/wiki/Statistics

```bash
--mc.world=path/to/world
```

### Usage ⚙

```bash
usage: minecraft-exporter [<flags>]

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --web.config.file=""       [EXPERIMENTAL] Path to configuration file that can enable TLS or authentication.
      --web.listen-address=":9150"
                                 Address to listen on for web interface and telemetry.
      --mc.config-path="config.yml"
                                 Path to YAML file with config.
      --mc.world="/minecraft/world"
                                 Path the to world folder
      --mc.rcon-address=":25575"
                                 Address of the Minecraft rcon.
      --mc.rcon-password=MC.RCON-PASSWORD
                                 Password of the Minecraft rcon.
      --mc.name-source="mojang"  How to retrieve names of players: offline, bukkit, mojang.
      --web.telemetry-path="/metrics"
                                 Path under which to expose metrics.
      --log.level=info           Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt        Output format of log messages. One of: [logfmt, json]
      --version                  Show application version.
```

### Config 🔧
You can override CLI flags using config file. By default, `config.yml` located in the current directory is used.
Path to config file can be changed using `--config-path` CLI flag.

| Key in config file | Equivalent CLI flag    | Description                                                                      |
| ---                | ---                    | ---                                                                              |
| `metrics-path`     | `--web.telemetry-path` | Path under which to expose metrics.                                              |
| `web-config`       | `--web.config.file`    | **[EXPERIMENTAL]** Path to configuration file that can enable TLS or authentication. |
| `listen-address`   | `--web.listen-address` | Address to listen on for web interface and telemetry.                            |
| `world-path`       | `--mc.world`           | Path the to world folder.                                                        |
| `rcon-address`     | `--mc.rcon-address`    | Address of the Minecraft rcon.                                                   |
| `rcon-password`    | `--mc.rcon-password`   | Password of the Minecraft rcon.                                                  |
| `name-source`      | `--mc.name-source`     | How to retrieve names of players: offline, bukkit, mojang.                       |
| `disabled-metrics` | -                      | Namespaced keys that used by metrics that should be disabled.                    |

#### Disabling metrics

To disable certain metrics, just add corresponding key to `disabled-metrics` section with `true` value in your config file.
You should use keys that used by Minecraft to store players' stats.

#### Example config

```yaml
disabled-metrics:
  minecraft:mined: true # Disable "minecraft_blocks_mined_total" metric
  minecraft:creeper: true # Disable all metrics related with creepers
  minecraft:custom: false # "false" values will be ignored, so this line does nothing
listen-address: ':9151' # Change address of web server. "--web.listen-address" will be ignored if this line is present here
```

### Collectors 📊
The exporter collects a number of statistics from the server:

```
# HELP minecraft_animals_bred The number of times the player bred two mobs.
# TYPE minecraft_animals_bred_total counter

# HELP minecraft_aviate_one_cm The total distance traveled by elytra.
# TYPE minecraft_aviate_one_cm_total counter

# HELP minecraft_bell_ring The number of times the player has rung a Bell.
# TYPE minecraft_bell_ring_total counter

# HELP minecraft_blocks_mined Statistic related to the number of blocks a player mined
# TYPE minecraft_blocks_mined_total counter

# HELP minecraft_boat_one_cm The total distance traveled by boats.
# TYPE minecraft_boat_one_cm_total counter

# HELP minecraft_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which minecraft was built.
# TYPE minecraft_build_info gauge

# HELP minecraft_clean_armor The number of dyed leather armors washed with a cauldron.
# TYPE minecraft_clean_armor_total counter

# HELP minecraft_clean_banner The number of banner patterns washed with a cauldron.
# TYPE minecraft_clean_banner_total counter

# HELP minecraft_clean_shulker_box The number of times the player has washed a Shulker Box with a cauldron.
# TYPE minecraft_clean_shulker_box_total counter

# HELP minecraft_climb_one_cm The total distance traveled up ladders or vines.
# TYPE minecraft_climb_one_cm_total counter

# HELP minecraft_crouch_one_cm The total distance walked while sneaking.
# TYPE minecraft_crouch_one_cm_total counter

# HELP minecraft_damage_absorbed The amount of damage the player has absorbed in tenths of 1♥.
# TYPE minecraft_damage_absorbed_total counter

# HELP minecraft_damage_blocked_by_shield The amount of damage the player has blocked with a shield in tenths of 1♥.
# TYPE minecraft_damage_blocked_by_shield_total counter

# HELP minecraft_damage_dealt The amount of damage the player has dealt in tenths 1♥. Includes only melee attacks.
# TYPE minecraft_damage_dealt_total counter

# HELP minecraft_damage_dealt_absorbed The amount of damage the player has dealt that were absorbed, in tenths of 1♥.
# TYPE minecraft_damage_dealt_absorbed_total counter

# HELP minecraft_damage_dealt_resisted The amount of damage the player has dealt that were resisted, in tenths of 1♥.
# TYPE minecraft_damage_dealt_resisted_total counter

# HELP minecraft_damage_resisted The amount of damage the player has resisted in tenths of 1♥.
# TYPE minecraft_damage_resisted_total counter

# HELP minecraft_damage_taken The amount of damage the player has taken in tenths of 1♥.
# TYPE minecraft_damage_taken_total counter

# HELP minecraft_deaths The number of times the player died.
# TYPE minecraft_deaths_total counter

# HELP minecraft_eat_cake_slice The number of cake slices eaten.
# TYPE minecraft_eat_cake_slice_total counter

# HELP minecraft_enchant_item The number of items enchanted.
# TYPE minecraft_enchant_item_total counter

# HELP minecraft_entities_killed Statistics related to the number of entities a player killed
# TYPE minecraft_entities_killed_total counter

# HELP minecraft_fall_one_cm The total distance fallen, excluding jumping. 
# TYPE minecraft_fall_one_cm_total counter

# HELP minecraft_fill_cauldron The number of times the player filled cauldrons with water buckets.
# TYPE minecraft_fill_cauldron_total counter

# HELP minecraft_fish_caught The number of fish caught.
# TYPE minecraft_fish_caught_total counter

# HELP minecraft_fly_one_cm Distance traveled upwards and forwards at the same time, while more than one block above the ground.
# TYPE minecraft_fly_one_cm_total counter

# HELP minecraft_horse_one_cm The total distance traveled by horses..
# TYPE minecraft_horse_one_cm_total counter

# HELP minecraft_inspect_dispenser The number of times interacted with dispensers.
# TYPE minecraft_inspect_dispenser_total counter

# HELP minecraft_inspect_dropper The number of times interacted with droppers.
# TYPE minecraft_inspect_dropper_total counter

# HELP minecraft_inspect_hopper The number of times interacted with hoppers.
# TYPE minecraft_inspect_hopper_total counter

# HELP minecraft_interact_with_anvil The number of times interacted with anvils.
# TYPE minecraft_interact_with_anvil_total counter

# HELP minecraft_interact_with_beacon The number of times interacted with beacons.
# TYPE minecraft_interact_with_beacon_total counter

# HELP minecraft_interact_with_blast_furnace The number of times interacted with blast furnaces
# TYPE minecraft_interact_with_blast_furnace_total counter

# HELP minecraft_interact_with_brewingstand The number of times interacted with brewing stands
# TYPE minecraft_interact_with_brewingstand_total counter

# HELP minecraft_interact_with_campfire The number of times interacted with campfires
# TYPE minecraft_interact_with_campfire_total counter

# HELP minecraft_interact_with_cartography_table The number of times interacted with cartography tables
# TYPE minecraft_interact_with_cartography_table_total counter

# HELP minecraft_interact_with_crafting_table The number of times interacted with crafting tables
# TYPE minecraft_interact_with_crafting_table_total counter

# HELP minecraft_interact_with_furnace The number of times interacted with furnaces
# TYPE minecraft_interact_with_furnace_total counter

# HELP minecraft_interact_with_grindstone The number of times interacted with grindstones
# TYPE minecraft_interact_with_grindstone_total counter

# HELP minecraft_interact_with_lectern The number of times interacted with lecterns
# TYPE minecraft_interact_with_lectern_total counter

# HELP minecraft_interact_with_loom The number of times interacted with looms
# TYPE minecraft_interact_with_loom_total counter

# HELP minecraft_interact_with_smithing_table The number of times interacted with smithing tables.
# TYPE minecraft_interact_with_smithing_table_total counter

# HELP minecraft_interact_with_smoker The number of times interacted with smokers.
# TYPE minecraft_interact_with_smoker_total counter

# HELP minecraft_interact_with_stonecutter The number of times interacted with stonecutters.
# TYPE minecraft_interact_with_stonecutter_total counter

# HELP minecraft_item_broken Statistics related to the number of items a player ran their durability negative
# TYPE minecraft_item_broken_total counter

# HELP minecraft_item_crafted Statistics related to the number of items crafted, smelted, etc.
# TYPE minecraft_item_crafted_total counter

# HELP minecraft_item_dropped Statistics related to the number of items that droped.
# TYPE minecraft_item_dropped_total counter

# HELP minecraft_item_picked_up Statistics related to the number of dropped items a player picked up
# TYPE minecraft_item_picked_up_total counter

# HELP minecraft_item_used Statistics related to the number of block or item used
# TYPE minecraft_item_used_total counter

# HELP minecraft_items_drop The number of items dropped.
# TYPE minecraft_items_drop_total counter

# HELP minecraft_jump       The total number of jumps performed.
# TYPE minecraft_jump_total counter

# HELP minecraft_killed_by Statistics related to the times of a player being killed by entities.
# TYPE minecraft_killed_by_total counter

# HELP minecraft_leave_game The number of times "Save and quit to title" has been clicked.
# TYPE minecraft_leave_game_total counter

# HELP minecraft_minecart_one_cm The total distance traveled by minecarts.
# TYPE minecraft_minecart_one_cm_total counter

# HELP minecraft_mob_kills The number of mobs the player killed.
# TYPE minecraft_mob_kills_total counter

# HELP minecraft_open_barrel The number of times the player has opened a Barrel.
# TYPE minecraft_open_barrel_total counter

# HELP minecraft_open_chest The number of times the player opened chests.
# TYPE minecraft_open_chest_total counter

# HELP minecraft_open_enderchest The number of times the player opened ender chests.
# TYPE minecraft_open_enderchest_total counter

# HELP minecraft_open_shulker_box The number of times the player has opened a Shulker Box.
# TYPE minecraft_open_shulker_box_total counter

# HELP minecraft_pig_one_cm The total distance traveled by pigs via saddles.
# TYPE minecraft_pig_one_cm_total counter

# HELP minecraft_play_noteblock The number of note blocks hit.
# TYPE minecraft_play_noteblock_total counter

# HELP minecraft_play_record The number of music discs played on a jukebox.
# TYPE minecraft_play_record_total counter

# HELP minecraft_play_time The total amount of time played. 
# TYPE minecraft_play_time_total counter

# HELP minecraft_player_advancements Number of completed advances of a player
# TYPE minecraft_player_advancements_total counter

# HELP minecraft_player_current_xp How much current XP a player has
# TYPE minecraft_player_current_xp_total counter

# HELP minecraft_player_food_level How much food the player currently has
# TYPE minecraft_player_food_level_total counter

# HELP minecraft_player_health How much Health the player currently has
# TYPE minecraft_player_health_total counter

# HELP minecraft_player_kills The number of players the player killed
# TYPE minecraft_player_kills_total counter

# HELP minecraft_player_online is 1 if player is online
# TYPE minecraft_player_online_total counter

# HELP minecraft_player_score The Score of the player
# TYPE minecraft_player_score_total counter

# HELP minecraft_player_xp_total How much total XP a player has
# TYPE minecraft_player_xp_total_total counter

# HELP minecraft_pot_flower The number of plants potted onto flower pots.
# TYPE minecraft_pot_flower_total counter

# HELP minecraft_raid_trigger The number of times the player has triggered a Raid.
# TYPE minecraft_raid_trigger_total counter

# HELP minecraft_raid_win The number of times the player has won a Raid.
# TYPE minecraft_raid_win_total counter

# HELP minecraft_sleep_in_bed The number of times the player has slept in a bed..
# TYPE minecraft_sleep_in_bed_total counter

# HELP minecraft_sneak_time The time the player has held down the sneak button.
# TYPE minecraft_sneak_time_total counter

# HELP minecraft_sprint_one_cm The total distance sprinted.
# TYPE minecraft_sprint_one_cm_total counter

# HELP minecraft_strider_one_cm The total distance traveled by striders via saddles.
# TYPE minecraft_strider_one_cm_total counter

# HELP minecraft_swim_one_cm The total distance covered with sprint-swimming..
# TYPE minecraft_swim_one_cm_total counter

# HELP minecraft_talked_to_villager The number of times interacted with villagers (opened the trading GUI).
# TYPE minecraft_talked_to_villager_total counter

# HELP minecraft_target_hit The number of times the player has shot a target block.
# TYPE minecraft_target_hit_total counter

# HELP minecraft_time_since_death The time since the player's last death.
# TYPE minecraft_time_since_death_total counter

# HELP minecraft_time_since_rest The time since the player's last rest. This is used to spawn phantoms.
# TYPE minecraft_time_since_rest_total counter

# HELP minecraft_total_world_time The total amount of time the world was opened.n.
# TYPE minecraft_total_world_time_total counter

# HELP minecraft_traded_with_villager The number of times traded with villagers.
# TYPE minecraft_traded_with_villager_total counter

# HELP minecraft_trigger_trapped_chest The number of times the player opened trapped chests.
# TYPE minecraft_trigger_trapped_chest_total counter

# HELP minecraft_tune_noteblock The number of times interacted with note blocks.
# TYPE minecraft_tune_noteblock_total counter

# HELP minecraft_use_cauldron The number of times the player took water from cauldrons with glass bottles.
# TYPE minecraft_use_cauldron_total counter

# HELP minecraft_walk_on_water_one_cm The distance covered while bobbing up and down over water.
# TYPE minecraft_walk_on_water_one_cm_total counter

# HELP minecraft_walk_one_cm The total distance walked.
# TYPE minecraft_walk_one_cm_total counter

# HELP minecraft_walk_under_water_one_cm The total distance you have walked underwater.
# TYPE minecraft_walk_under_water_one_cm_total counter
```

### Libraries & Tools 🔥

- https://github.com/Jeffail/gabs
- https://github.com/alecthomas/kingpin
- https://github.com/Tnze/go-mc
- https://github.com/prometheus/exporter-toolkit
- https://github.com/goreleaser

### Legal Disclaimer 👮

This project is not affiliated with Mojang Studios, XBox Game Studios, Double Eleven or the Minecraft brand.

"Minecraft" is a trademark of Mojang Synergies AB.

Other trademarks referenced herein are property of their respective owners.