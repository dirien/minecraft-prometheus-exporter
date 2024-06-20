###### ⚠️ Statistics are present only in Java Edition. Bedrock Edition has no equivalent of statistics in-game ⚠️

# Minecraft Exporter for Prometheus

![minecraft-exporter](https://dirien.github.io/minecraft-prometheus-exporter/img/minecraft-exporter.jpg)

![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)
![Minecraft](https://img.shields.io/badge/Minecraft-62B47A?style=for-the-badge&logo=Minecraft&logoColor=white)
![Grafana](https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=Grafana&logoColor=white)
![Docker](https://img.shields.io/badge/docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Chaingard Images Static](https://img.shields.io/badge/chainguard_image_static-4445E7?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Helm](https://img.shields.io/badge/helm-0F1689?style=for-the-badge&logo=helm&logoColor=white)

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/minecraft-exporter&style=for-the-badge)](https://artifacthub.io/packages/search?repo=minecraft-exporter)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/minecraft-exporter-image&style=for-the-badge)](https://artifacthub.io/packages/search?repo=minecraft-exporter-image)

![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/dirien/minecraft-prometheus-exporter/Build%20Binary/main?logo=github&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/dirien/minecraft-prometheus-exporter?style=for-the-badge)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/dirien/minecraft-prometheus-exporter?style=for-the-badge)

This is a Prometheus Minecraft exporter, created as part of the [minectl 🗺](https://github.com/dirien/minectl) project.

It collects metrics from different sources of the game

- RCON
- NBT Files
- Advancement file
- Player stats file

#### Getting started

Short getting started guide with systemd unit file on ubuntu linux

##### Create Systemd unit file

```bash
tee /etc/systemd/system/minecraft-exporter.service <<EOF
[Unit]
Description=Minecraft Exporter
Wants=network-online.target
After=network-online.target
[Service]
User=minecraft_exporter
Group=minecraft_exporter
Type=simple
ExecStart=/usr/local/bin/minecraft-exporter \
  --mc.rcon-password=<rcon password if needed>
[Install]
WantedBy=multi-user.target
EOF
```

#### Install Minecraft Exporter

```
MINECRAFT_EXPORTER_VERSION=0.6.1
curl -sSL https://github.com/dirien/minecraft-prometheus-exporter/releases/download/v$MINECRAFT_EXPORTER_VERSION/minecraft-exporter_$MINECRAFT_EXPORTER_VERSION.linux-$ARCH.tar.gz | tar -xz
cp minecraft-exporter /usr/local/bin
chown minecraft_exporter:minecraft_exporter /usr/local/bin/minecraft-exporter
systemctl start minecraft-exporter.service
systemctl enable minecraft-exporter.service
```

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

or

```bash
export MC_RCON_ADDRESS=":25575"
export MC_RCON_PASSWORD=<rcon-password>
```

#### Files

To access the files, take care that the `minecraft-exporter` binary is started with the flag

Details for the specific stats can be found here -> https://minecraft.fandom.com/wiki/Statistics

```bash
--mc.world=path/to/world
```

or

```bash
export MC_WORLD=path/to/world
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
      --web.telemetry-path="/metrics"  
                                 Path under which to expose metrics.
      --web.listen-address=":9150"  
                                 Address to listen on for web interface and telemetry.
      --web.disable-exporter-metrics  
                                 Disabling collection of exporter metrics (like go_*)
      --mc.config-path="config.yml"  
                                 Path to YAML file with config.
      --mc.world="/minecraft/world"  
                                 Path the to world folder
      --mc.rcon-address=":25575"  
                                 Address of the Minecraft rcon.
      --mc.rcon-password=MC.RCON-PASSWORD  
                                 Password of the Minecraft rcon.
      --mc.name-source="mojang"  How to retrieve names of players: offline, bukkit, mojang.
      --mc.mod-server-stats=MC.MOD-SERVER-STATS  
                                 Additional server stats for papermc, purpurmc, forge, or fabric.
      --log.level=info           Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt        Output format of log messages. One of: [logfmt, json]
      --version                  Show application version.
```

### Config 🔧

You can override CLI flags using config file. By default, `config.yml` located in the current directory is used. Path to
config file can be changed using `--config-path` CLI flag.

| Key in config file         | Equivalent CLI flag              | ENV variable                  | Description                                                           |
|----------------------------|----------------------------------|-------------------------------|-----------------------------------------------------------------------|
| `metrics-path`             | `--web.telemetry-path`           | WEB_TELEMETRY_PATH            | Path under which to expose metrics.                                   |
| `listen-address`           | `--web.listen-address`           | WEB_LISTEN_ADDRESS            | Address to listen on for web interface and telemetry.                 |
| `disable-exporter-metrics` | `--web.disable-exporter-metrics` | WEB_DISABLED_EXPORTER_METRICS | Disabling collection of exporter metrics (like go_*)                  |
| `web-config`               | `--mc.config-path`               | MC_CONFIG_PATH                | Path to YAML file with config for the mc variables                    |
| `world-path`               | `--mc.world`                     | MC_WORLD                      | Path to the world folder.                                             |
| `rcon-address`             | `--mc.rcon-address`              | MC_RCON_ADDRESS               | Address for the Minecraft RCON.                                       |
| `rcon-password`            | `--mc.rcon-password`             | MC_RCON_PASSWORD              | Password for the Minecraft RCON.                                      |
| `name-source`              | `--mc.name-source`               | MC_NAME_SOURCE                | How to retrieve names of players: offline, bukkit, mojang.            |
| `mod-server-stats`         | `--mc.mod-server-stats`          | MC_MOD_SERVER_STATS           | Set server for additional stats (papermc, purpurmc, forge, or fabric) |

#### Disable exporter metrics

With the flag `--web.disable-exporter-metrics` you can disable collection of exporter metrics (like go_*). This is
useful if you want just see the minecraft metrics and not the exporter metrics.

#### Disabling metrics

To disable certain metrics, just add corresponding key to `disabled-metrics` section with `true` value in your config
file. You should use keys that used by Minecraft to store players' stats.

#### Example config (--mc.config-path/MC_CONFIG_PATH)

```yaml
disabled-metrics:
  minecraft:mined: true # Disable "minecraft_blocks_mined_total" metric
  minecraft:creeper: true # Disable all metrics related with creepers
  minecraft:custom: false # "false" values will be ignored, so this line does nothing
listen-address: ':9151' # Change address of web server. "--web.listen-address" will be ignored if this line is present here
```

### Known Issues ⚠️️

`minecraft-exporter` is heavily dependent on the username of the player. If you are using mods like [EssentialsX](https://essentialsx.net/)
and hide or change the default grouping of players, then `minecraft-exporter` will may not work properly.

> If you encounter any issues, please let me know by opening an GitHub issue in this repository.

### Collectors 📊

The exporter collects a number of statistics from the server (with example of the labels):

```bash
# HELP minecraft_animals_breded_total The number of times the player bred two mobs
# TYPE minecraft_animals_breded_total counter
minecraft_animals_breded_total{player="ediri"} 0
# HELP minecraft_bells_ringed_total TThe number of times the player rang a bell
# TYPE minecraft_bells_ringed_total counter
minecraft_bells_ringed_total{player="ediri"} 0
# HELP minecraft_blocks_mined_total Statistic related to the number of blocks a player mined
# TYPE minecraft_blocks_mined_total counter
minecraft_blocks_mined_total{block="birch_log",namespace="minecraft",player="ediri"} 2
minecraft_blocks_mined_total{block="clay",namespace="minecraft",player="ediri"} 1
minecraft_blocks_mined_total{block="coal_ore",namespace="minecraft",player="ediri"} 16
minecraft_blocks_mined_total{block="crafting_table",namespace="minecraft",player="ediri"} 2
minecraft_blocks_mined_total{block="dirt",namespace="minecraft",player="ediri"} 19
minecraft_blocks_mined_total{block="grass",namespace="minecraft",player="ediri"} 38
minecraft_blocks_mined_total{block="grass_block",namespace="minecraft",player="ediri"} 18
minecraft_blocks_mined_total{block="oak_leaves",namespace="minecraft",player="ediri"} 5
minecraft_blocks_mined_total{block="oak_log",namespace="minecraft",player="ediri"} 11
minecraft_blocks_mined_total{block="oak_planks",namespace="minecraft",player="ediri"} 1
minecraft_blocks_mined_total{block="oxeye_daisy",namespace="minecraft",player="ediri"} 1
minecraft_blocks_mined_total{block="seagrass",namespace="minecraft",player="ediri"} 5
minecraft_blocks_mined_total{block="spruce_leaves",namespace="minecraft",player="ediri"} 11
minecraft_blocks_mined_total{block="spruce_log",namespace="minecraft",player="ediri"} 13
minecraft_blocks_mined_total{block="stone",namespace="minecraft",player="ediri"} 2
minecraft_blocks_mined_total{block="tall_grass",namespace="minecraft",player="ediri"} 1
minecraft_blocks_mined_total{block="tall_seagrass",namespace="minecraft",player="ediri"} 5
minecraft_blocks_mined_total{block="vine",namespace="minecraft",player="ediri"} 2
minecraft_blocks_mined_total{block="wall_torch",namespace="minecraft",player="ediri"} 1
# HELP minecraft_cake_slices_eaten_total The number of times the player ate cake
# TYPE minecraft_cake_slices_eaten_total counter
minecraft_cake_slices_eaten_total{player="ediri"} 0
# HELP minecraft_cleaned_armors_total The number of times the player washed dyed leather armor with a cauldron
# TYPE minecraft_cleaned_armors_total counter
minecraft_cleaned_armors_total{player="ediri"} 0
# HELP minecraft_cleaned_banner_total The number of times the player washed a banner with a cauldron
# TYPE minecraft_cleaned_banner_total counter
minecraft_cleaned_banner_total{player="ediri"} 0
# HELP minecraft_damage_dealt_total The amount of damage the player has dealt of different types (in tenths of 1♥)
# TYPE minecraft_damage_dealt_total counter
minecraft_damage_dealt_total{player="ediri",type="absorbed"} 0
minecraft_damage_dealt_total{player="ediri",type="hit"} 610
minecraft_damage_dealt_total{player="ediri",type="resisted"} 0
# HELP minecraft_damage_received_total The amount of damage the player has taken of different types (in tenths of 1♥)
# TYPE minecraft_damage_received_total counter
minecraft_damage_received_total{player="ediri",type="absorbed"} 0
minecraft_damage_received_total{player="ediri",type="blocked_by_shield"} 0
minecraft_damage_received_total{player="ediri",type="resisted"} 0
minecraft_damage_received_total{player="ediri",type="taken"} 2295
# HELP minecraft_deaths_total The number of times the player died
# TYPE minecraft_deaths_total counter
minecraft_deaths_total{player="ediri"} 9
# HELP minecraft_entities_killed_total Statistics related to the number of entities a player killed
# TYPE minecraft_entities_killed_total counter
minecraft_entities_killed_total{entity="pig",player="ediri"} 2
# HELP minecraft_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which minecraft_exporter was built.
# TYPE minecraft_exporter_build_info gauge
minecraft_exporter_build_info{branch="",goversion="go1.16.5",revision="",version=""} 1
# HELP minecraft_filled_cauldrons_total The number of times the player filled a cauldron with a water bucket
# TYPE minecraft_filled_cauldrons_total counter
minecraft_filled_cauldrons_total{player="ediri"} 0
# HELP minecraft_fishs_caught_total The number of times the player caught fish
# TYPE minecraft_fishs_caught_total counter
minecraft_fishs_caught_total{player="ediri"} 0
# HELP minecraft_games_left_total The number of times the player clicked "Save and quit to title"
# TYPE minecraft_games_left_total counter
minecraft_games_left_total{player="ediri"} 12
# HELP minecraft_inspected_total The number of times the player inspected a dispenser, hopper or dropper
# TYPE minecraft_inspected_total counter
minecraft_inspected_total{entity="dispenser",player="ediri"} 0
minecraft_inspected_total{entity="dropper",player="ediri"} 0
minecraft_inspected_total{entity="hopper",player="ediri"} 0
# HELP minecraft_interactions_total The number of times the player interacted with different entities
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
# HELP minecraft_item_actions_total Statistics related to items and the number of times they were used, picked up, dropped or broken
# TYPE minecraft_item_actions_total counter
minecraft_item_actions_total{action="broken",namespace="minecraft",entity="wooden_axe",player="ediri"} 1
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="crafting_table",player="ediri"} 3
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="oak_planks",player="ediri"} 32
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="oak_wood",player="ediri"} 6
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="spruce_planks",player="ediri"} 28
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="stick",player="ediri"} 32
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="torch",player="ediri"} 8
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="wooden_axe",player="ediri"} 1
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="wooden_pickaxe",player="ediri"} 3
minecraft_item_actions_total{action="crafted",namespace="minecraft",entity="wooden_shovel",player="ediri"} 1
minecraft_item_actions_total{action="dropped",namespace="minecraft",entity="porkchop",player="ediri"} 4
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="birch_log",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="coal",player="ediri"} 15
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="crafting_table",player="ediri"} 2
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="dirt",player="ediri"} 33
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="oak_log",player="ediri"} 11
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="oak_planks",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="oak_sapling",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="oxeye_daisy",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="porkchop",player="ediri"} 4
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="spruce_log",player="ediri"} 13
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="torch",player="ediri"} 1
minecraft_item_actions_total{action="picked_up",namespace="minecraft",entity="wheat_seeds",player="ediri"} 6
minecraft_item_actions_total{action="used",namespace="minecraft",entity="crafting_table",player="ediri"} 3
minecraft_item_actions_total{action="used",namespace="minecraft",entity="oak_planks",player="ediri"} 1
minecraft_item_actions_total{action="used",namespace="minecraft",entity="oxeye_daisy",player="ediri"} 1
minecraft_item_actions_total{action="used",namespace="minecraft",entity="torch",player="ediri"} 2
minecraft_item_actions_total{action="used",namespace="minecraft",entity="wooden_axe",player="ediri"} 52
minecraft_item_actions_total{action="used",namespace="minecraft",entity="wooden_pickaxe",player="ediri"} 26
# HELP minecraft_items_dropped_total The number of items the player dropped
# TYPE minecraft_items_dropped_total counter
minecraft_items_dropped_total{player="ediri"} 1
# HELP minecraft_items_enchanted_total The number of items the player enchanted
# TYPE minecraft_items_enchanted_total counter
minecraft_items_enchanted_total{player="ediri"} 0
# HELP minecraft_jumps_total The number of times the player jumped
# TYPE minecraft_jumps_total counter
minecraft_jumps_total{player="ediri"} 406
# HELP minecraft_killed_by_total Statistics related to the number of times a player was killed by entities
# TYPE minecraft_killed_by_total counter
minecraft_killed_by_total{entity="skeleton",namespace="minecraft",player="ediri"} 1
minecraft_killed_by_total{entity="zombie",namespace="minecraft",player="ediri"} 7
# HELP minecraft_mobs_killed_total The number of mobs the player killed
# TYPE minecraft_mobs_killed_total counter
minecraft_mobs_killed_total{player="ediri"} 2
# HELP minecraft_movement_meters_total The total distance the player traveled by different methods (ladders, boats, swimming etc.)
# TYPE minecraft_movement_meters_total untyped
minecraft_movement_meters_total{means="aviate",player="ediri"} 0
minecraft_movement_meters_total{means="boat",player="ediri"} 0
minecraft_movement_meters_total{means="climb",player="ediri"} 0.16
minecraft_movement_meters_total{means="crouch",player="ediri"} 0
minecraft_movement_meters_total{means="fall",player="ediri"} 58.49
minecraft_movement_meters_total{means="fly",player="ediri"} 55.14
minecraft_movement_meters_total{means="horse",player="ediri"} 0
minecraft_movement_meters_total{means="minecart",player="ediri"} 0
minecraft_movement_meters_total{means="pig",player="ediri"} 0
minecraft_movement_meters_total{means="sprint",player="ediri"} 0.17
minecraft_movement_meters_total{means="strider",player="ediri"} 0
minecraft_movement_meters_total{means="swim",player="ediri"} 0
minecraft_movement_meters_total{means="walk",player="ediri"} 1411.7
minecraft_movement_meters_total{means="walk_on_water",player="ediri"} 168.45
minecraft_movement_meters_total{means="walk_under_water",player="ediri"} 156.92
# HELP minecraft_noteblocks_played_total The number of times the player hit a note block
# TYPE minecraft_noteblocks_played_total counter
minecraft_noteblocks_played_total{player="ediri"} 0
# HELP minecraft_noteblocks_tuned_total The number of times the player tuned a note block
# TYPE minecraft_noteblocks_tuned_total counter
minecraft_noteblocks_tuned_total{player="ediri"} 0
# HELP minecraft_opened_barrels_total The number of times the player opened a barrel
# TYPE minecraft_opened_barrels_total counter
minecraft_opened_barrels_total{player="ediri"} 0
# HELP minecraft_opened_chests_total The number of times the player opened a chest
# TYPE minecraft_opened_chests_total counter
minecraft_opened_chests_total{player="ediri"} 0
# HELP minecraft_opened_enderchests_total The number of times the player opened an ender chest
# TYPE minecraft_opened_enderchests_total counter
minecraft_opened_enderchests_total{player="ediri"} 0
# HELP minecraft_play_time_ticks_total The number of ticks the player has played
# TYPE minecraft_play_time_ticks_total counter
minecraft_play_time_ticks_total{player="ediri"} 0
# HELP minecraft_player_online_total Players currently online (1 if player is online)
# TYPE minecraft_player_online_total counter
minecraft_player_online_total{player="ediri"} 1
# HELP minecraft_player_stat_total Statistic related to the player: xp, current_xp, food_level, health, score, advancements
# TYPE minecraft_player_stat_total gauge
minecraft_player_stat_total{player="ediri",stat="advancements"} 53
minecraft_player_stat_total{player="ediri",stat="current_xp"} 0
minecraft_player_stat_total{player="ediri",stat="food_level"} 20
minecraft_player_stat_total{player="ediri",stat="health"} 20
minecraft_player_stat_total{player="ediri",stat="score"} 1
minecraft_player_stat_total{player="ediri",stat="xp"} 1
# HELP minecraft_players_killed_total The number of times the player killed a player
# TYPE minecraft_players_killed_total counter
minecraft_players_killed_total{player="ediri"} 0
# HELP minecraft_pots_flowered_total The number of times the player planted a plant in a flower pot
# TYPE minecraft_pots_flowered_total counter
minecraft_pots_flowered_total{player="ediri"} 0
# HELP minecraft_records_played_total The number of times the player played a music disc on a jukebox
# TYPE minecraft_records_played_total counter
minecraft_records_played_total{player="ediri"} 0
# HELP minecraft_shulker_boxes_cleaned_total The number of times the player washed a shulker box with a cauldron
# TYPE minecraft_shulker_boxes_cleaned_total counter
minecraft_shulker_boxes_cleaned_total{player="ediri"} 0
# HELP minecraft_shulker_boxes_opened_total The number of times the player opened a shulker box
# TYPE minecraft_shulker_boxes_opened_total counter
minecraft_shulker_boxes_opened_total{player="ediri"} 0
# HELP minecraft_sleep_in_bed_ticks_total The number of times the player slept in a bed
# TYPE minecraft_sleep_in_bed_ticks_total counter
minecraft_sleep_in_bed_ticks_total{player="ediri"} 0
# HELP minecraft_sneak_time_ticks_total The number of ticks the player has spent sneaking
# TYPE minecraft_sneak_time_ticks_total counter
minecraft_sneak_time_ticks_total{player="ediri"} 0
# HELP minecraft_talked_to_villagers_total The number of times the player spoke with a villager (opened the trading GUI)
# TYPE minecraft_talked_to_villagers_total counter
minecraft_talked_to_villagers_total{player="ediri"} 0
# HELP minecraft_targets_hit_total The number of times the player shot a target block
# TYPE minecraft_targets_hit_total counter
minecraft_targets_hit_total{player="ediri"} 0
# HELP minecraft_time_since_death_ticks_total The number of ticks since the player's last death
# TYPE minecraft_time_since_death_ticks_total counter
minecraft_time_since_death_ticks_total{player="ediri"} 24217
# HELP minecraft_time_since_rest_ticks_total The number of ticks since the player's last rest (used to spawn phantoms)
# TYPE minecraft_time_since_rest_ticks_total counter
minecraft_time_since_rest_ticks_total{player="ediri"} 24238
# HELP minecraft_total_world_time_ticks_total The number of ticks the player has been in the world
# TYPE minecraft_total_world_time_ticks_total counter
minecraft_total_world_time_ticks_total{player="ediri"} 159695
# HELP minecraft_traded_with_villagers_total The number of times the player traded with a villager
# TYPE minecraft_traded_with_villagers_total counter
minecraft_traded_with_villagers_total{player="ediri"} 0
# HELP minecraft_triggered_raids_total The number of times the player triggered a raid
# TYPE minecraft_triggered_raids_total counter
minecraft_triggered_raids_total{player="ediri"} 0
# HELP minecraft_triggered_trapped_chests_total The number of times the player opened a trapped chest
# TYPE minecraft_triggered_trapped_chests_total counter
minecraft_triggered_trapped_chests_total{player="ediri"} 0
# HELP minecraft_used_cauldrons_total The number of times the player took water from cauldrons with glass bottles
# TYPE minecraft_used_cauldrons_total counter
minecraft_used_cauldrons_total{player="ediri"} 0
# HELP minecraft_won_raids_total The number of times the player won a raid
# TYPE minecraft_won_raids_total counter
minecraft_won_raids_total{player="ediri"} 0
```

#### Metrics for specific mods

To export the metrics for a specific mod, you can set the flag: `--mc.mod-server-stats=forge|papermc|fabric` or use the
environment variable `MC_MOD_SERVER_STATS`.

**Note:** Fabric requires the [Fabric TPS](https://modrinth.com/mod/fabric-tps) mod.

#### Forge

`minecraft-exporter` exports mean TPS and mean tick time for every dimension and of course the overall mean TPS and mean
tick time.

Additionally, you will get the total numbers of active entities on the server per dimension.

```bash
# HELP minecraft_ticktime_total The overall mean tick time in the server
# TYPE minecraft_ticktime_total counter
minecraft_ticktime_total 0.38
# HELP minecraft_tps_total The overall mean ticks per second in the server
# TYPE minecraft_tps_total counter
minecraft_tps_total 20
# HELP minecraft_dimension_ticktime_total The mean tick time in a certain dimension
# TYPE minecraft_dimension_ticktime_total counter
minecraft_dimension_ticktime_total{dimension="overworld",namespace="minecraft"} 0.37
minecraft_dimension_ticktime_total{dimension="the_end",namespace="minecraft"} 0.002
minecraft_dimension_ticktime_total{dimension="the_nether",namespace="minecraft"} 0.004
# HELP minecraft_dimension_tps_total The number of ticks per second in a certain dimension
# TYPE minecraft_dimension_tps_total counter
minecraft_dimension_tps_total{dimension="overworld",namespace="minecraft"} 20
minecraft_dimension_tps_total{dimension="the_end",namespace="minecraft"} 20
minecraft_dimension_tps_total{dimension="the_nether",namespace="minecraft"} 20
# HELP minecraft_active_entity_total The number and type of an active entity on the server
# TYPE minecraft_active_entity_total counter
minecraft_active_entity_total{entity="bat",namespace="minecraft"} 15
minecraft_active_entity_total{entity="bee",namespace="minecraft"} 7
minecraft_active_entity_total{entity="chest_minecart",namespace="minecraft"} 15
minecraft_active_entity_total{entity="chicken",namespace="minecraft"} 8
minecraft_active_entity_total{entity="cod",namespace="minecraft"} 9
minecraft_active_entity_total{entity="cow",namespace="minecraft"} 12
minecraft_active_entity_total{entity="creeper",namespace="minecraft"} 12
minecraft_active_entity_total{entity="dolphin",namespace="minecraft"} 1
minecraft_active_entity_total{entity="drowned",namespace="minecraft"} 1
minecraft_active_entity_total{entity="enderman",namespace="minecraft"} 6
minecraft_active_entity_total{entity="falling_block",namespace="minecraft"} 1
minecraft_active_entity_total{entity="item",namespace="minecraft"} 13
minecraft_active_entity_total{entity="pig",namespace="minecraft"} 12
minecraft_active_entity_total{entity="pufferfish",namespace="minecraft"} 6
minecraft_active_entity_total{entity="rabbit",namespace="minecraft"} 1
minecraft_active_entity_total{entity="sheep",namespace="minecraft"} 17
minecraft_active_entity_total{entity="skeleton",namespace="minecraft"} 28
minecraft_active_entity_total{entity="spider",namespace="minecraft"} 6
minecraft_active_entity_total{entity="squid",namespace="minecraft"} 5
minecraft_active_entity_total{entity="tropical_fish",namespace="minecraft"} 8
minecraft_active_entity_total{entity="zombie",namespace="minecraft"} 19
```

#### PaperMC

`minecraft-exporter` exports the TPS from the last 1m, 5m, and 15m as histogram.

```bash
# HELP minecraft_tps_total_bucket The number of ticks per second in PaperMC
# TYPE minecraft_tps_total_bucket histogram
minecraft_tps_total_bucket_bucket{le="1"} 20
minecraft_tps_total_bucket_bucket{le="5"} 20
minecraft_tps_total_bucket_bucket{le="15"} 20
minecraft_tps_total_bucket_bucket{le="+Inf"} 3
minecraft_tps_total_bucket_sum 60
minecraft_tps_total_bucket_count 3
```

#### PurpurMC

To get in `PurpurMC` TPS in a form `minecraft-exporter` can read it you need to enable `overrideTpsCommand` in the 
`spark` configuration file. Check the [spark docs](https://spark.lucko.me/docs/Configuration#overridetpscommand) for
details on how to do this.

`minecraft-exporter` exports the TPS from the last 5s (here written in 0.08 minute), 1m, 5m, and 15m as histogram.

```bash
# HELP minecraft_tps_total_bucket The number of ticks per second in PaperMC
# TYPE minecraft_tps_total_bucket histogram
minecraft_tps_total_bucket_bucket{le="0.08"} 20
minecraft_tps_total_bucket_bucket{le="1"} 20
minecraft_tps_total_bucket_bucket{le="5"} 20
minecraft_tps_total_bucket_bucket{le="15"} 20
minecraft_tps_total_bucket_bucket{le="+Inf"} 4
minecraft_tps_total_bucket_sum 80
minecraft_tps_total_bucket_count 4
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

### Stargazers over time 🌟

[![Stargazers over time](https://starchart.cc/dirien/minecraft-prometheus-exporter.svg)](https://starchart.cc/dirien/minecraft-prometheus-exporter) 
