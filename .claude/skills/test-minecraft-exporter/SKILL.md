---
name: test-minecraft-exporter
description: End-to-end docker-compose test harness for the minecraft-prometheus-exporter. Spins up an itzg/minecraft-server:latest vanilla server alongside a freshly built exporter and walks the user through verifying metrics in a live browser/Prometheus session. Use this skill whenever the user wants to "test the exporter end-to-end", "verify metrics against a real Minecraft server", "spin up a local MC server", "run the docker compose test", scrape /metrics from a vanilla MC server, or sanity-check a code change in this repo before opening a PR — including cases where the user just says "let me see if it works" or "run the exporter against MC" without naming the compose file.
---

# Test the minecraft-prometheus-exporter end-to-end

This is a runbook for booting `docker-compose.test.yml` and confirming the exporter scrapes a real Minecraft server. Run these steps in order from the repo root.

## Files this skill depends on

Both live at the repo root and are checked into git:

- `docker-compose.test.yml` — defines the `minecraft` (itzg/minecraft-server:latest, vanilla, RCON on, offline mode) and `exporter` services, plus a shared `mc-data` named volume so the exporter can read the world dir.
- `Dockerfile.dev` — multi-stage `golang:1.26-alpine` → `chainguard/static` build of the exporter from source. Compose builds this image automatically with `--build`.

If either file is missing, do not improvise — flag it to the user. The compose setup has UID/volume specifics that are easy to get wrong.

## 1. Bring the stack up

```bash
docker compose -f docker-compose.test.yml up -d --build
```

First boot takes ~90 seconds because the vanilla MC server has to download the JAR (resolved from `VERSION: LATEST`) and generate a world. The compose file has a `service_healthy` dependency on the MC container's `mc-health` healthcheck, so the exporter will only start once MC is ready — but `docker compose up` returns earlier than that. Tell the user this is expected; subsequent boots are much faster.

## 2. Endpoints

Once both containers are up, tell the user:

- **Exporter metrics:** http://localhost:9150/metrics
- **Exporter landing page:** http://localhost:9150/
- **Minecraft server (Java client):** `localhost:25565`

## 3. Identify the MC version

`VERSION: LATEST` means "whatever itzg's image considers latest at pull time" — useful but not self-documenting. Show the user which version was resolved so they know which client to use:

```bash
docker exec mc-exporter-test-server bash -c 'ls /data/*.jar'
```

The jar filename embeds the version (e.g. `minecraft_server.26.1.2.jar`).

## 4. Connect a client

The server runs with `ONLINE_MODE=FALSE`, so any Java Edition client at the matching version can connect with any username — no Mojang/Microsoft auth needed. Just add a server pointed at `localhost:25565`.

Until a player connects at least once, MC doesn't create `world/playerdata/`, `world/stats/`, or `world/advancements/`. With an empty server, `/metrics` will only show `minecraft_exporter_build_info`. That's expected, not a bug.

## 5. Force a flush and inspect metrics

After the user has logged in and done a few things in-game (mine a block, walk around, jump), push MC to flush its in-memory stats:

```bash
docker exec mc-exporter-test-server rcon-cli save-all
```

Then re-scrape:

```bash
curl -s http://localhost:9150/metrics | grep '^minecraft_'
```

You should now see populated series — `minecraft_blocks_mined_total`, `minecraft_jumps_total`, `minecraft_movement_meters_total{means="walk"}`, `minecraft_player_online{player="..."}`, etc.

Quick sanity-check RCON too:

```bash
docker exec mc-exporter-test-server rcon-cli list
```

## 6. The label-name quirk

The compose sets `MC_NAME_SOURCE=offline`. This means:

- **RCON-derived series** (`minecraft_player_online`) carry the readable player name (e.g. `player="_diri"`), because the RCON `list` command returns names directly.
- **File-derived series** (everything in `world/stats/`, `world/playerdata/`, `world/advancements/`) carry the raw UUID (e.g. `player="8d34c913-…"`), because the exporter only resolves names when `name-source` is `mojang` (calls playerdb.co) or `bukkit` (reads `bukkit.lastKnownName` from NBT).

In vanilla + offline mode neither alternative works: the offline UUID isn't registered with Mojang, and vanilla doesn't write the Bukkit name field. So `offline` (UUID-as-label) is the only honest choice for this test stack. If the user wants names everywhere they need a real online server or a Paper/Spigot/Purpur server. Don't try to "fix" this by switching `MC_NAME_SOURCE` — it will silently make things worse.

## 7. Tear down

```bash
docker compose -f docker-compose.test.yml down -v
```

The `-v` flag wipes the `mc-data` volume, so the next run starts from a fresh world. Drop `-v` if the user wants to keep the world (and their player progress) between runs — first boot is the slow one, subsequent boots reuse the existing world and skip generation.

## Troubleshooting

**`permission denied` on `/mc/world/playerdata` in the exporter logs:**
Confirm the `exporter` service in `docker-compose.test.yml` still has `user: "1000:1000"`. The itzg image writes the world as UID 1000, and the chainguard/static base used by the exporter defaults to UID 65532, so without the UID override the read-only mount is unreadable. This is the single most likely thing to break if someone refactors the compose file.

**Exporter starts before MC is ready and crashes on RCON dial:**
Shouldn't happen — the compose file uses `depends_on: minecraft: condition: service_healthy`. If it does, check whether someone removed the healthcheck or the `condition:` block.

**`/metrics` only ever shows `minecraft_exporter_build_info`:**
No player has joined the server yet. Connect at least once, do something in-game, then `rcon-cli save-all` and re-scrape. The exporter has no synthetic mode — empty world means empty metrics.

**MC won't pull a new version / boot loops:**
`VERSION: LATEST` can occasionally lag behind a brand-new Minecraft release while itzg's image catches up. Pin to a specific version (e.g. `VERSION: "1.21.4"`) in the compose file as a workaround.
