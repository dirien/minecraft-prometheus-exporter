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

#### API

Due to restrictions of the API from mojang `https://api.mojang.com/user/profiles/[uuid]/names` we switched to the
project of [Electroid](https://github.com/Electroid/mojang-api)

Mojang, the developers of Minecraft, provides multiple APIs for websites and servers to fetch identity information about
users. Requests do not accept authentication tokens, however they are heavily rate limited and fragmented among several
endpoints. The purpose of this project is to package several of the most commonly used APIs into a single GET request
with no rate limiting and no need for client-side caching.

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

You can override CLI flags using config file. By default, `config.yml` located in the current directory is used. Path to
config file can be changed using `--config-path` CLI flag.

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

To disable certain metrics, just add corresponding key to `disabled-metrics` section with `true` value in your config
file. You should use keys that used by Minecraft to store players' stats.

#### Example config

```yaml
disabled-metrics:
  minecraft:mined: true # Disable "minecraft_blocks_mined_total" metric
  minecraft:creeper: true # Disable all metrics related with creepers
  minecraft:custom: false # "false" values will be ignored, so this line does nothing
listen-address: ':9151' # Change address of web server. "--web.listen-address" will be ignored if this line is present here
```

### Collectors 📊

The exporter collects a number of statistics from the server (with example of the labels):

```bash
# HELP minecraft_animals_bred_total The number of times the player bred two mobs.
# TYPE minecraft_animals_bred_total counter
minecraft_animals_bred_total{player="ediri"} 0
# HELP minecraft_bell_ring_total The number of times the player has rung a Bell.
# TYPE minecraft_bell_ring_total counter
minecraft_bell_ring_total{player="ediri"} 0
# HELP minecraft_blocks_mined_total Statistic related to the number of blocks a player mined
# TYPE minecraft_blocks_mined_total counter
minecraft_blocks_mined_total{block="birch_log",player="ediri"} 2
minecraft_blocks_mined_total{block="clay",player="ediri"} 1
minecraft_blocks_mined_total{block="coal_ore",player="ediri"} 16
minecraft_blocks_mined_total{block="crafting_table",player="ediri"} 2
minecraft_blocks_mined_total{block="dirt",player="ediri"} 19
minecraft_blocks_mined_total{block="grass",player="ediri"} 38
minecraft_blocks_mined_total{block="grass_block",player="ediri"} 18
minecraft_blocks_mined_total{block="oak_leaves",player="ediri"} 5
minecraft_blocks_mined_total{block="oak_log",player="ediri"} 11
minecraft_blocks_mined_total{block="oak_planks",player="ediri"} 1
minecraft_blocks_mined_total{block="oxeye_daisy",player="ediri"} 1
minecraft_blocks_mined_total{block="seagrass",player="ediri"} 5
minecraft_blocks_mined_total{block="spruce_leaves",player="ediri"} 11
minecraft_blocks_mined_total{block="spruce_log",player="ediri"} 13
minecraft_blocks_mined_total{block="stone",player="ediri"} 2
minecraft_blocks_mined_total{block="tall_grass",player="ediri"} 1
minecraft_blocks_mined_total{block="tall_seagrass",player="ediri"} 5
minecraft_blocks_mined_total{block="vine",player="ediri"} 2
minecraft_blocks_mined_total{block="wall_torch",player="ediri"} 1
# HELP minecraft_clean_armor_total The number of dyed leather armors washed with a cauldron.
# TYPE minecraft_clean_armor_total counter
minecraft_clean_armor_total{player="ediri"} 0
# HELP minecraft_clean_banner_total The number of banner patterns washed with a cauldron.
# TYPE minecraft_clean_banner_total counter
minecraft_clean_banner_total{player="ediri"} 0
# HELP minecraft_clean_shulker_box_total The number of times the player has washed a Shulker Box with a cauldron.
# TYPE minecraft_clean_shulker_box_total counter
minecraft_clean_shulker_box_total{player="ediri"} 0
# HELP minecraft_damage_total The amount of damage the player has handled from different types in tenths of 1♥.
# TYPE minecraft_damage_total counter
minecraft_damage_total{player="ediri",type="absorbed"} 0
minecraft_damage_total{player="ediri",type="blocked_by_shield"} 0
minecraft_damage_total{player="ediri",type="dealt"} 610
minecraft_damage_total{player="ediri",type="dealt_absorbed"} 0
minecraft_damage_total{player="ediri",type="dealt_resisted"} 0
minecraft_damage_total{player="ediri",type="resisted"} 0
minecraft_damage_total{player="ediri",type="taken"} 2295
# HELP minecraft_deaths_total The number of times the player died.
# TYPE minecraft_deaths_total counter
minecraft_deaths_total{player="ediri"} 9
# HELP minecraft_eat_cake_slice_total The number of cake slices eaten.
# TYPE minecraft_eat_cake_slice_total counter
minecraft_eat_cake_slice_total{player="ediri"} 0
# HELP minecraft_enchant_item_total The number of items enchanted.
# TYPE minecraft_enchant_item_total counter
minecraft_enchant_item_total{player="ediri"} 0
# HELP minecraft_entities_killed_total Statistics related to the number of entities a player killed
# TYPE minecraft_entities_killed_total counter
minecraft_entities_killed_total{entity="pig",player="ediri"} 2
# HELP minecraft_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which minecraft_exporter was built.
# TYPE minecraft_exporter_build_info gauge
minecraft_exporter_build_info{branch="",goversion="go1.16.5",revision="",version=""} 1
# HELP minecraft_fill_cauldron_total The number of times the player filled cauldrons with water buckets.
# TYPE minecraft_fill_cauldron_total counter
minecraft_fill_cauldron_total{player="ediri"} 0
# HELP minecraft_fish_caught_total The number of fish caught.
# TYPE minecraft_fish_caught_total counter
minecraft_fish_caught_total{player="ediri"} 0
# HELP minecraft_inspected_total The number of times inspected a dispenser, hopper or dropper.
# TYPE minecraft_inspected_total counter
minecraft_inspected_total{entity="dispenser",player="ediri"} 0
minecraft_inspected_total{entity="dropper",player="ediri"} 0
minecraft_inspected_total{entity="hopper",player="ediri"} 0
# HELP minecraft_interactions_total The number of times interacted with different entities
# TYPE minecraft_interactions_total counter
minecraft_interactions_total{entity="anvil",player="ediri"} 0
minecraft_interactions_total{entity="beacon",player="ediri"} 0
minecraft_interactions_total{entity="blast_furnace",player="ediri"} 0
minecraft_interactions_total{entity="brewingstand",player="ediri"} 0
minecraft_interactions_total{entity="campfire",player="ediri"} 0
minecraft_interactions_total{entity="cartography_table",player="ediri"} 0
minecraft_interactions_total{entity="crafting_table",player="ediri"} 10
minecraft_interactions_total{entity="furnace",player="ediri"} 0
minecraft_interactions_total{entity="grindston",player="ediri"} 0
minecraft_interactions_total{entity="lectern",player="ediri"} 0
minecraft_interactions_total{entity="loom",player="ediri"} 0
minecraft_interactions_total{entity="smithing_table",player="ediri"} 0
minecraft_interactions_total{entity="smoker",player="ediri"} 0
minecraft_interactions_total{entity="stonecutter",player="ediri"} 0
# HELP minecraft_item_actions_total Statistics related to the number of items and their actions: used, picked up, dropped, broken
# TYPE minecraft_item_actions_total counter
minecraft_item_actions_total{action="broken",entity="wooden_axe",player="ediri"} 1
minecraft_item_actions_total{action="crafted",entity="crafting_table",player="ediri"} 3
minecraft_item_actions_total{action="crafted",entity="oak_planks",player="ediri"} 32
minecraft_item_actions_total{action="crafted",entity="oak_wood",player="ediri"} 6
minecraft_item_actions_total{action="crafted",entity="spruce_planks",player="ediri"} 28
minecraft_item_actions_total{action="crafted",entity="stick",player="ediri"} 32
minecraft_item_actions_total{action="crafted",entity="torch",player="ediri"} 8
minecraft_item_actions_total{action="crafted",entity="wooden_axe",player="ediri"} 1
minecraft_item_actions_total{action="crafted",entity="wooden_pickaxe",player="ediri"} 3
minecraft_item_actions_total{action="crafted",entity="wooden_shovel",player="ediri"} 1
minecraft_item_actions_total{action="dropped",entity="porkchop",player="ediri"} 4
minecraft_item_actions_total{action="picked_up",entity="birch_log",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",entity="coal",player="ediri"} 15
minecraft_item_actions_total{action="picked_up",entity="crafting_table",player="ediri"} 2
minecraft_item_actions_total{action="picked_up",entity="dirt",player="ediri"} 33
minecraft_item_actions_total{action="picked_up",entity="oak_log",player="ediri"} 11
minecraft_item_actions_total{action="picked_up",entity="oak_planks",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",entity="oak_sapling",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",entity="oxeye_daisy",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",entity="porkchop",player="ediri"} 4
minecraft_item_actions_total{action="picked_up",entity="spruce_log",player="ediri"} 13
minecraft_item_actions_total{action="picked_up",entity="torch",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",entity="wheat_seeds",player="ediri"} 6
minecraft_item_actions_total{action="used",entity="crafting_table",player="ediri"} 3
minecraft_item_actions_total{action="used",entity="oak_planks",player="ediri"} 1
minecraft_item_actions_total{action="used",entity="oxeye_daisy",player="ediri"} 1
minecraft_item_actions_total{action="used",entity="torch",player="ediri"} 2
minecraft_item_actions_total{action="used",entity="wooden_axe",player="ediri"} 52
minecraft_item_actions_total{action="used",entity="wooden_pickaxe",player="ediri"} 26
# HELP minecraft_items_drop_total The number of items dropped.
# TYPE minecraft_items_drop_total counter
minecraft_items_drop_total{player="ediri"} 1
# HELP minecraft_jump_total     The total number of jumps performed.
# TYPE minecraft_jump_total counter
minecraft_jump_total{player="ediri"} 406
# HELP minecraft_killed_by_total Statistics related to the times of a player being killed by entities.
# TYPE minecraft_killed_by_total counter
minecraft_killed_by_total{entity="skeleton",player="ediri"} 1
minecraft_killed_by_total{entity="zombie",player="ediri"} 7
# HELP minecraft_leave_game_total The number of times "Save and quit to title" has been clicked.
# TYPE minecraft_leave_game_total counter
minecraft_leave_game_total{player="ediri"} 11
# HELP minecraft_mob_kills_total The number of mobs the player killed.
# TYPE minecraft_mob_kills_total counter
minecraft_mob_kills_total{player="ediri"} 2
# HELP minecraft_movement_meter_total The total distance traveled with different entities (ladders, boats, etc.)
# TYPE minecraft_movement_meter_total untyped
minecraft_movement_meter_total{means="aviate",player="ediri"} 0
minecraft_movement_meter_total{means="boat",player="ediri"} 0
minecraft_movement_meter_total{means="climb",player="ediri"} 0.16
minecraft_movement_meter_total{means="crouch",player="ediri"} 0
minecraft_movement_meter_total{means="fall",player="ediri"} 58.49
minecraft_movement_meter_total{means="fly",player="ediri"} 55.14
minecraft_movement_meter_total{means="horse",player="ediri"} 0
minecraft_movement_meter_total{means="minecart",player="ediri"} 0
minecraft_movement_meter_total{means="pig",player="ediri"} 0
minecraft_movement_meter_total{means="sprint",player="ediri"} 0.17
minecraft_movement_meter_total{means="strider",player="ediri"} 0
minecraft_movement_meter_total{means="swim",player="ediri"} 0
minecraft_movement_meter_total{means="walk",player="ediri"} 1411.7
minecraft_movement_meter_total{means="walk_on_water",player="ediri"} 168.45
minecraft_movement_meter_total{means="walk_under_water",player="ediri"} 156.92
# HELP minecraft_open_barrel_total The number of times the player has opened a Barrel.
# TYPE minecraft_open_barrel_total counter
minecraft_open_barrel_total{player="ediri"} 0
# HELP minecraft_open_chest_total The number of times the player opened chests.
# TYPE minecraft_open_chest_total counter
minecraft_open_chest_total{player="ediri"} 0
# HELP minecraft_open_enderchest_total The number of times the player opened ender chests.
# TYPE minecraft_open_enderchest_total counter
minecraft_open_enderchest_total{player="ediri"} 0
# HELP minecraft_open_shulker_box_total The number of times the player has opened a Shulker Box.
# TYPE minecraft_open_shulker_box_total counter
minecraft_open_shulker_box_total{player="ediri"} 0
# HELP minecraft_play_noteblock_total The number of note blocks hit.
# TYPE minecraft_play_noteblock_total counter
minecraft_play_noteblock_total{player="ediri"} 0
# HELP minecraft_play_record_total The number of music discs played on a jukebox.
# TYPE minecraft_play_record_total counter
minecraft_play_record_total{player="ediri"} 0
# HELP minecraft_play_time_total The total amount of time played.
# TYPE minecraft_play_time_total counter
minecraft_play_time_total{player="ediri"} 0
# HELP minecraft_player_kills_total The number of players the player killed
# TYPE minecraft_player_kills_total counter
minecraft_player_kills_total{player="ediri"} 0
# HELP minecraft_player_stat_total Different stats of the player: xp, current_xp, food_level, health, score, advancements
# TYPE minecraft_player_stat_total counter
minecraft_player_stat_total{player="ediri",stat="advancements"} 53
minecraft_player_stat_total{player="ediri",stat="current_xp"} 0
minecraft_player_stat_total{player="ediri",stat="food_level"} 20
minecraft_player_stat_total{player="ediri",stat="health"} 20
minecraft_player_stat_total{player="ediri",stat="score"} 1
minecraft_player_stat_total{player="ediri",stat="xp"} 1
# HELP minecraft_pot_flower_total The number of plants potted onto flower pots.
# TYPE minecraft_pot_flower_total counter
minecraft_pot_flower_total{player="ediri"} 0
# HELP minecraft_raid_trigger_total The number of times the player has triggered a Raid.
# TYPE minecraft_raid_trigger_total counter
minecraft_raid_trigger_total{player="ediri"} 0
# HELP minecraft_raid_win_total The number of times the player has won a Raid.
# TYPE minecraft_raid_win_total counter
minecraft_raid_win_total{player="ediri"} 0
# HELP minecraft_sleep_in_bed_total The number of times the player has slept in a bed..
# TYPE minecraft_sleep_in_bed_total counter
minecraft_sleep_in_bed_total{player="ediri"} 0
# HELP minecraft_sneak_time_total The time the player has held down the sneak button.
# TYPE minecraft_sneak_time_total counter
minecraft_sneak_time_total{player="ediri"} 0
# HELP minecraft_talked_to_villager_total The number of times interacted with villagers (opened the trading GUI).
# TYPE minecraft_talked_to_villager_total counter
minecraft_talked_to_villager_total{player="ediri"} 0
# HELP minecraft_target_hit_total The number of times the player has shot a target block.
# TYPE minecraft_target_hit_total counter
minecraft_target_hit_total{player="ediri"} 0
# HELP minecraft_time_since_death_total The time since the player's last death.
# TYPE minecraft_time_since_death_total counter
minecraft_time_since_death_total{player="ediri"} 3373
# HELP minecraft_time_since_rest_total The time since the player's last rest. This is used to spawn phantoms.
# TYPE minecraft_time_since_rest_total counter
minecraft_time_since_rest_total{player="ediri"} 3394
# HELP minecraft_total_world_time_total The total amount of time the world was opened.n.
# TYPE minecraft_total_world_time_total counter
minecraft_total_world_time_total{player="ediri"} 138851
# HELP minecraft_traded_with_villager_total The number of times traded with villagers.
# TYPE minecraft_traded_with_villager_total counter
minecraft_traded_with_villager_total{player="ediri"} 0
# HELP minecraft_trigger_trapped_chest_total The number of times the player opened trapped chests.
# TYPE minecraft_trigger_trapped_chest_total counter
minecraft_trigger_trapped_chest_total{player="ediri"} 0
# HELP minecraft_tune_noteblock_total The number of times interacted with note blocks.
# TYPE minecraft_tune_noteblock_total counter
minecraft_tune_noteblock_total{player="ediri"} 0
# HELP minecraft_use_cauldron_total The number of times the player took water from cauldrons with glass bottles.
# TYPE minecraft_use_cauldron_total counter
minecraft_use_cauldron_total{player="ediri"} 0
```

### Libraries & Tools 🔥

- https://github.com/Jeffail/gabs
- https://github.com/alecthomas/kingpin
- https://github.com/Tnze/go-mc
- https://github.com/prometheus/exporter-toolkit
- https://github.com/goreleaser
- https://github.com/Electroid/mojang-api

### Legal Disclaimer 👮

This project is not affiliated with Mojang Studios, XBox Game Studios, Double Eleven or the Minecraft brand.

"Minecraft" is a trademark of Mojang Synergies AB.

Other trademarks referenced herein are property of their respective owners.