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

To enable rcon on your minecraft server add the following to the `properties` field in your `minectl` manifest:

```yaml
    ...
    edition: java
    properties: |
      broadcast-rcon-to-ops=false
      rcon.port=25575
      rcon.password=<your password>
      enable-rcon=true
    ...
```

See [server config examples](https://github.com/dirien/minectl#server-config-) for details and usage of the minectl 🗺
manifest file.

Alternatively you need to edit your server.properties on your server.

Be sure that following flags are set for the `minecraft-prometheus-exporter` binary:

```bash
--mc.rcon-address=":25575"  
--mc.rcon-password=<rcon-password>
```

#### Files

To access the files, take care that the `minecraft-prometheus-exporter` binary is started with the flag

Details for the specific stats can be found here -> https://minecraft.fandom.com/wiki/Statistics

```bash
--mc.world=path/to/world
```

### Usage ⚙

```bash
usage: minecraft-prometheus-exporter [<flags>]

Flags:
  -h, --help                Show context-sensitive help (also try --help-long and --help-man).
      --web.config.file=""  [EXPERIMENTAL] Path to configuration file that can enable TLS or authentication.
      --web.listen-address=":9150"  
                            Address to listen on for web interface and telemetry.
      --mc.world="/minecraft/world"  
                            Path the to world folder
      --mc.rcon-address=":25575"  
                            Address of the Minecraft rcon.
      --mc.rcon-password=MC.RCON-PASSWORD  
                            Password of the Minecraft rcon.
      --web.telemetry-path="/metrics"  
                            Path under which to expose metrics.
      --log.level=info      Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt   Output format of log messages. One of: [logfmt, json]
      --version             Show application version.
```

### Collectors 📊
The exporter collects a number of statistics from the server:

```
# HELP minecraft_prometheus_exporter_animals_bred The number of times the player bred two mobs.
# TYPE minecraft_prometheus_exporter_animals_bred counter

# HELP minecraft_prometheus_exporter_aviate_one_cm The total distance traveled by elytra.
# TYPE minecraft_prometheus_exporter_aviate_one_cm counter

# HELP minecraft_prometheus_exporter_bell_ring The number of times the player has rung a Bell.
# TYPE minecraft_prometheus_exporter_bell_ring counter

# HELP minecraft_prometheus_exporter_blocks_mined Statistic related to the number of blocks a player mined
# TYPE minecraft_prometheus_exporter_blocks_mined counter

# HELP minecraft_prometheus_exporter_boat_one_cm The total distance traveled by boats.
# TYPE minecraft_prometheus_exporter_boat_one_cm counter

# HELP minecraft_prometheus_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which minecraft_prometheus_exporter was built.
# TYPE minecraft_prometheus_exporter_build_info gauge

# HELP minecraft_prometheus_exporter_clean_armor The number of dyed leather armors washed with a cauldron.
# TYPE minecraft_prometheus_exporter_clean_armor counter

# HELP minecraft_prometheus_exporter_clean_banner The number of banner patterns washed with a cauldron.
# TYPE minecraft_prometheus_exporter_clean_banner counter

# HELP minecraft_prometheus_exporter_clean_shulker_box The number of times the player has washed a Shulker Box with a cauldron.
# TYPE minecraft_prometheus_exporter_clean_shulker_box counter

# HELP minecraft_prometheus_exporter_climb_one_cm The total distance traveled up ladders or vines.
# TYPE minecraft_prometheus_exporter_climb_one_cm counter

# HELP minecraft_prometheus_exporter_crouch_one_cm The total distance walked while sneaking.
# TYPE minecraft_prometheus_exporter_crouch_one_cm counter

# HELP minecraft_prometheus_exporter_damage_absorbed The amount of damage the player has absorbed in tenths of 1♥.
# TYPE minecraft_prometheus_exporter_damage_absorbed counter

# HELP minecraft_prometheus_exporter_damage_blocked_by_shield The amount of damage the player has blocked with a shield in tenths of 1♥.
# TYPE minecraft_prometheus_exporter_damage_blocked_by_shield counter

# HELP minecraft_prometheus_exporter_damage_dealt The amount of damage the player has dealt in tenths 1♥. Includes only melee attacks.
# TYPE minecraft_prometheus_exporter_damage_dealt counter

# HELP minecraft_prometheus_exporter_damage_dealt_absorbed The amount of damage the player has dealt that were absorbed, in tenths of 1♥.
# TYPE minecraft_prometheus_exporter_damage_dealt_absorbed counter

# HELP minecraft_prometheus_exporter_damage_dealt_resisted The amount of damage the player has dealt that were resisted, in tenths of 1♥.
# TYPE minecraft_prometheus_exporter_damage_dealt_resisted counter

# HELP minecraft_prometheus_exporter_damage_resisted The amount of damage the player has resisted in tenths of 1♥.
# TYPE minecraft_prometheus_exporter_damage_resisted counter

# HELP minecraft_prometheus_exporter_damage_taken The amount of damage the player has taken in tenths of 1♥.
# TYPE minecraft_prometheus_exporter_damage_taken counter

# HELP minecraft_prometheus_exporter_deaths The number of times the player died.
# TYPE minecraft_prometheus_exporter_deaths counter

# HELP minecraft_prometheus_exporter_eat_cake_slice The number of cake slices eaten.
# TYPE minecraft_prometheus_exporter_eat_cake_slice counter

# HELP minecraft_prometheus_exporter_enchant_item The number of items enchanted.
# TYPE minecraft_prometheus_exporter_enchant_item counter

# HELP minecraft_prometheus_exporter_entities_killed Statistics related to the number of entities a player killed
# TYPE minecraft_prometheus_exporter_entities_killed counter

# HELP minecraft_prometheus_exporter_fall_one_cm The total distance fallen, excluding jumping. 
# TYPE minecraft_prometheus_exporter_fall_one_cm counter

# HELP minecraft_prometheus_exporter_fill_cauldron The number of times the player filled cauldrons with water buckets.
# TYPE minecraft_prometheus_exporter_fill_cauldron counter

# HELP minecraft_prometheus_exporter_fish_caught The number of fish caught.
# TYPE minecraft_prometheus_exporter_fish_caught counter

# HELP minecraft_prometheus_exporter_fly_one_cm Distance traveled upwards and forwards at the same time, while more than one block above the ground.
# TYPE minecraft_prometheus_exporter_fly_one_cm counter

# HELP minecraft_prometheus_exporter_horse_one_cm The total distance traveled by horses..
# TYPE minecraft_prometheus_exporter_horse_one_cm counter

# HELP minecraft_prometheus_exporter_inspect_dispenser The number of times interacted with dispensers.
# TYPE minecraft_prometheus_exporter_inspect_dispenser counter

# HELP minecraft_prometheus_exporter_inspect_dropper The number of times interacted with droppers.
# TYPE minecraft_prometheus_exporter_inspect_dropper counter

# HELP minecraft_prometheus_exporter_inspect_hopper The number of times interacted with hoppers.
# TYPE minecraft_prometheus_exporter_inspect_hopper counter

# HELP minecraft_prometheus_exporter_interact_with_anvil The number of times interacted with anvils.
# TYPE minecraft_prometheus_exporter_interact_with_anvil counter

# HELP minecraft_prometheus_exporter_interact_with_beacon The number of times interacted with beacons.
# TYPE minecraft_prometheus_exporter_interact_with_beacon counter

# HELP minecraft_prometheus_exporter_interact_with_blast_furnace The number of times interacted with blast furnaces
# TYPE minecraft_prometheus_exporter_interact_with_blast_furnace counter

# HELP minecraft_prometheus_exporter_interact_with_brewingstand The number of times interacted with brewing stands
# TYPE minecraft_prometheus_exporter_interact_with_brewingstand counter

# HELP minecraft_prometheus_exporter_interact_with_campfire The number of times interacted with campfires
# TYPE minecraft_prometheus_exporter_interact_with_campfire counter

# HELP minecraft_prometheus_exporter_interact_with_cartography_table The number of times interacted with cartography tables
# TYPE minecraft_prometheus_exporter_interact_with_cartography_table counter

# HELP minecraft_prometheus_exporter_interact_with_crafting_table The number of times interacted with crafting tables
# TYPE minecraft_prometheus_exporter_interact_with_crafting_table counter

# HELP minecraft_prometheus_exporter_interact_with_furnace The number of times interacted with furnaces
# TYPE minecraft_prometheus_exporter_interact_with_furnace counter

# HELP minecraft_prometheus_exporter_interact_with_grindstone The number of times interacted with grindstones
# TYPE minecraft_prometheus_exporter_interact_with_grindstone counter

# HELP minecraft_prometheus_exporter_interact_with_lectern The number of times interacted with lecterns
# TYPE minecraft_prometheus_exporter_interact_with_lectern counter

# HELP minecraft_prometheus_exporter_interact_with_loom The number of times interacted with looms
# TYPE minecraft_prometheus_exporter_interact_with_loom counter

# HELP minecraft_prometheus_exporter_interact_with_smithing_table The number of times interacted with smithing tables.
# TYPE minecraft_prometheus_exporter_interact_with_smithing_table counter

# HELP minecraft_prometheus_exporter_interact_with_smoker The number of times interacted with smokers.
# TYPE minecraft_prometheus_exporter_interact_with_smoker counter

# HELP minecraft_prometheus_exporter_interact_with_stonecutter The number of times interacted with stonecutters.
# TYPE minecraft_prometheus_exporter_interact_with_stonecutter counter

# HELP minecraft_prometheus_exporter_item_broken Statistics related to the number of items a player ran their durability negative
# TYPE minecraft_prometheus_exporter_item_broken counter

# HELP minecraft_prometheus_exporter_item_crafted Statistics related to the number of items crafted, smelted, etc.
# TYPE minecraft_prometheus_exporter_item_crafted counter

# HELP minecraft_prometheus_exporter_item_dropped Statistics related to the number of items that droped.
# TYPE minecraft_prometheus_exporter_item_dropped counter

# HELP minecraft_prometheus_exporter_item_picked_up Statistics related to the number of dropped items a player picked up
# TYPE minecraft_prometheus_exporter_item_picked_up counter

# HELP minecraft_prometheus_exporter_item_used Statistics related to the number of block or item used
# TYPE minecraft_prometheus_exporter_item_used counter

# HELP minecraft_prometheus_exporter_items_drop The number of items dropped.
# TYPE minecraft_prometheus_exporter_items_drop counter

# HELP minecraft_prometheus_exporter_jump       The total number of jumps performed.
# TYPE minecraft_prometheus_exporter_jump counter

# HELP minecraft_prometheus_exporter_killed_by Statistics related to the times of a player being killed by entities.
# TYPE minecraft_prometheus_exporter_killed_by counter

# HELP minecraft_prometheus_exporter_leave_game The number of times "Save and quit to title" has been clicked.
# TYPE minecraft_prometheus_exporter_leave_game counter

# HELP minecraft_prometheus_exporter_minecart_one_cm The total distance traveled by minecarts.
# TYPE minecraft_prometheus_exporter_minecart_one_cm counter

# HELP minecraft_prometheus_exporter_mob_kills The number of mobs the player killed.
# TYPE minecraft_prometheus_exporter_mob_kills counter

# HELP minecraft_prometheus_exporter_open_barrel The number of times the player has opened a Barrel.
# TYPE minecraft_prometheus_exporter_open_barrel counter

# HELP minecraft_prometheus_exporter_open_chest The number of times the player opened chests.
# TYPE minecraft_prometheus_exporter_open_chest counter

# HELP minecraft_prometheus_exporter_open_enderchest The number of times the player opened ender chests.
# TYPE minecraft_prometheus_exporter_open_enderchest counter

# HELP minecraft_prometheus_exporter_open_shulker_box The number of times the player has opened a Shulker Box.
# TYPE minecraft_prometheus_exporter_open_shulker_box counter

# HELP minecraft_prometheus_exporter_pig_one_cm The total distance traveled by pigs via saddles.
# TYPE minecraft_prometheus_exporter_pig_one_cm counter

# HELP minecraft_prometheus_exporter_play_noteblock The number of note blocks hit.
# TYPE minecraft_prometheus_exporter_play_noteblock counter

# HELP minecraft_prometheus_exporter_play_record The number of music discs played on a jukebox.
# TYPE minecraft_prometheus_exporter_play_record counter

# HELP minecraft_prometheus_exporter_play_time The total amount of time played. 
# TYPE minecraft_prometheus_exporter_play_time counter

# HELP minecraft_prometheus_exporter_player_advancements Number of completed advances of a player
# TYPE minecraft_prometheus_exporter_player_advancements counter

# HELP minecraft_prometheus_exporter_player_current_xp How much current XP a player has
# TYPE minecraft_prometheus_exporter_player_current_xp counter

# HELP minecraft_prometheus_exporter_player_food_level How much food the player currently has
# TYPE minecraft_prometheus_exporter_player_food_level counter

# HELP minecraft_prometheus_exporter_player_health How much Health the player currently has
# TYPE minecraft_prometheus_exporter_player_health counter

# HELP minecraft_prometheus_exporter_player_kills The number of players the player killed
# TYPE minecraft_prometheus_exporter_player_kills counter

# HELP minecraft_prometheus_exporter_player_online is 1 if player is online
# TYPE minecraft_prometheus_exporter_player_online counter

# HELP minecraft_prometheus_exporter_player_score The Score of the player
# TYPE minecraft_prometheus_exporter_player_score counter

# HELP minecraft_prometheus_exporter_player_xp_total How much total XP a player has
# TYPE minecraft_prometheus_exporter_player_xp_total counter

# HELP minecraft_prometheus_exporter_pot_flower The number of plants potted onto flower pots.
# TYPE minecraft_prometheus_exporter_pot_flower counter

# HELP minecraft_prometheus_exporter_raid_trigger The number of times the player has triggered a Raid.
# TYPE minecraft_prometheus_exporter_raid_trigger counter

# HELP minecraft_prometheus_exporter_raid_win The number of times the player has won a Raid.
# TYPE minecraft_prometheus_exporter_raid_win counter

# HELP minecraft_prometheus_exporter_sleep_in_bed The number of times the player has slept in a bed..
# TYPE minecraft_prometheus_exporter_sleep_in_bed counter

# HELP minecraft_prometheus_exporter_sneak_time The time the player has held down the sneak button.
# TYPE minecraft_prometheus_exporter_sneak_time counter

# HELP minecraft_prometheus_exporter_sprint_one_cm The total distance sprinted.
# TYPE minecraft_prometheus_exporter_sprint_one_cm counter

# HELP minecraft_prometheus_exporter_strider_one_cm The total distance traveled by striders via saddles.
# TYPE minecraft_prometheus_exporter_strider_one_cm counter

# HELP minecraft_prometheus_exporter_swim_one_cm The total distance covered with sprint-swimming..
# TYPE minecraft_prometheus_exporter_swim_one_cm counter

# HELP minecraft_prometheus_exporter_talked_to_villager The number of times interacted with villagers (opened the trading GUI).
# TYPE minecraft_prometheus_exporter_talked_to_villager counter

# HELP minecraft_prometheus_exporter_target_hit The number of times the player has shot a target block.
# TYPE minecraft_prometheus_exporter_target_hit counter

# HELP minecraft_prometheus_exporter_time_since_death The time since the player's last death.
# TYPE minecraft_prometheus_exporter_time_since_death counter

# HELP minecraft_prometheus_exporter_time_since_rest The time since the player's last rest. This is used to spawn phantoms.
# TYPE minecraft_prometheus_exporter_time_since_rest counter

# HELP minecraft_prometheus_exporter_total_world_time The total amount of time the world was opened.n.
# TYPE minecraft_prometheus_exporter_total_world_time counter

# HELP minecraft_prometheus_exporter_traded_with_villager The number of times traded with villagers.
# TYPE minecraft_prometheus_exporter_traded_with_villager counter

# HELP minecraft_prometheus_exporter_trigger_trapped_chest The number of times the player opened trapped chests.
# TYPE minecraft_prometheus_exporter_trigger_trapped_chest counter

# HELP minecraft_prometheus_exporter_tune_noteblock The number of times interacted with note blocks.
# TYPE minecraft_prometheus_exporter_tune_noteblock counter

# HELP minecraft_prometheus_exporter_use_cauldron The number of times the player took water from cauldrons with glass bottles.
# TYPE minecraft_prometheus_exporter_use_cauldron counter

# HELP minecraft_prometheus_exporter_walk_on_water_one_cm The distance covered while bobbing up and down over water.
# TYPE minecraft_prometheus_exporter_walk_on_water_one_cm counter

# HELP minecraft_prometheus_exporter_walk_one_cm The total distance walked.
# TYPE minecraft_prometheus_exporter_walk_one_cm counter

# HELP minecraft_prometheus_exporter_walk_under_water_one_cm The total distance you have walked underwater.
# TYPE minecraft_prometheus_exporter_walk_under_water_one_cm counter
```

### Libraries & Tools 🔥

- https://github.com/Jeffail/gabs
- https://github.com/gorcon/rcon
- https://github.com/alecthomas/kingpin
- https://github.com/Tnze/go-mc
- https://github.com/prometheus/exporter-toolkit
- https://github.com/goreleaser

### Legal Disclaimer 👮

This project is not affiliated with Mojang Studios, XBox Game Studios, Double Eleven or the Minecraft brand.

"Minecraft" is a trademark of Mojang Synergies AB.

Other trademarks referenced herein are property of their respective owners.