# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands

```bash
# Build
go build .

# Run locally (requires RCON-enabled Minecraft server)
go run . --mc.rcon-address="localhost:25575" --mc.rcon-password="<password>" --mc.world="/path/to/world"

# Lint
golangci-lint run --timeout=5m

# Update dependencies
go get -u ./... && go mod tidy
```

## Architecture

This is a Prometheus exporter for Minecraft Java Edition servers. It collects metrics from multiple sources:

- **RCON**: Remote console protocol for real-time server data (player list, TPS for modded servers)
- **NBT Files**: Player data files (`/playerdata/*.dat`) for XP, health, food level
- **Stats JSON**: Player statistics (`/stats/*.json`) for blocks mined, items used, deaths, etc.
- **Advancements JSON**: Player advancements (`/advancements/*.json`)

### Key Files

- `main.go` → Entry point, delegates to `cmd/minecraftexporter`
- `cmd/minecraftexporter/main.go` → HTTP server setup, Prometheus registry, signal handling
- `pkg/exporter/exporter.go` → Core exporter implementing Prometheus Collector interface (~1000 lines)
- `pkg/config/config.go` → CLI flags (kingpin) and YAML config parsing
- `pkg/template/template.go` → Landing page HTML template

### Metric Collection Flow

The exporter implements `prometheus.Collector` with:
- `Describe()` → Declares all metric descriptors
- `Collect()` → Called on each scrape, executes:
  1. `getPlayerList()` → RCON `list` command → `minecraft_player_online` gauge
  2. `getPlayerStats()` → Reads NBT/JSON files → 50+ player stat metrics
  3. `getServerStats()` → RCON commands for Forge/PaperMC/Fabric → TPS/ticktime gauges

### Player Name Resolution

The `--mc.name-source` flag controls how player UUIDs are mapped to usernames:
- `mojang` (default): Uses playerdb.co API (requires online-mode server)
- `bukkit`: Reads from NBT `bukkit.lastKnownName` field
- `offline`: Uses raw UUID as player label

### Modded Server Support

Set `--mc.mod-server-stats` for additional metrics:
- `forge`: TPS, ticktime per dimension, entity counts via `forge tps` and `forge entity list`
- `papermc`: TPS histogram via `tps` command
- `purpurmc`: Extended TPS histogram (5s, 1m, 5m, 15m buckets)
- `fabric`: Similar to Forge (requires Fabric TPS mod)

## Metric Conventions

- Counters use `_total` suffix and `prometheus.CounterValue`
- Gauges (point-in-time values like TPS, health, online status) use `prometheus.GaugeValue`
- Ticktime metrics include `_ms` suffix for milliseconds
- Player stats use the `player` label for per-player breakdown
