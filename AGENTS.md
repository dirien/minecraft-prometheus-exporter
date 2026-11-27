# AGENTS.md

Guidance for AI agents working in this repository. Humans should read `README.md` and `CLAUDE.md` instead.

## Project

Prometheus exporter for Minecraft Java Edition servers. Single Go binary that scrapes RCON + reads world files (`playerdata/`, `stats/`, `advancements/`) and exposes them at `/metrics`.

## Build / lint / run

Authoritative source: `CLAUDE.md`. The short version:

| Task | Command |
|------|---------|
| Build | `go build .` |
| Lint | `golangci-lint run --timeout=5m` |
| Update deps | `go get -u ./... && go mod tidy` |
| Run locally | `go run . --mc.rcon-address="host:25575" --mc.rcon-password="..." --mc.world="/path/to/world"` |

If you edit those commands, also edit `CLAUDE.md` — they must stay in sync.

## Flag parsing — important

Flags are defined via [`alecthomas/kong`](https://github.com/alecthomas/kong) struct tags on `Config` in `pkg/config/config.go`. **Do not reach for `alecthomas/kingpin/v2`** — it was removed during the kong migration. Reasons that matter:

- Every flag has an `env:"..."` tag, so `WEB_LISTEN_ADDRESS`, `MC_*`, `LOG_*` etc. are all honored. This was the fix for issue #1061; don't regress it by re-introducing `prometheus/exporter-toolkit/web/kingpinflag.AddFlags`, which does **not** bind env vars.
- The toolkit's `*web.FlagConfig` is assembled by `(*Config).WebFlagConfig()` from the kong-parsed fields; the `--web.systemd-socket` field is forced to `false` on non-Linux to mirror upstream behavior.
- `--log.level` / `--log.format` are declared on `Config` directly (not via `prometheus/common/promslog/flag`, which is kingpin-bound). `cmd/minecraftexporter/main.go` constructs `*promslog.Config` from those fields manually.

YAML overlay via `--mc.config-path` (default `config.yml`) still works — `Config.LoadFile()` unmarshals yaml-tagged fields onto the kong-populated struct, so YAML wins over flags/env when present.

## Layout

| Path | Purpose |
|------|---------|
| `main.go` | Entry point → delegates to `cmd/minecraftexporter` |
| `cmd/minecraftexporter/main.go` | HTTP server, Prometheus registry, signal handling, kong parsing |
| `pkg/config/config.go` | Kong-tagged `Config`, YAML overlay, `WebFlagConfig()` builder |
| `pkg/exporter/exporter.go` | `prometheus.Collector` — RCON + NBT + JSON readers (~1000 lines) |
| `pkg/template/` | Landing page HTML |
| `charts/` | Helm chart |
| `Dockerfile` | Release image (assumes pre-built binary, used by goreleaser) |
| `Dockerfile.dev` | Source-build image used only by `docker-compose.test.yml` |

## End-to-end testing

There is a working docker-compose test harness for live verification against a vanilla Minecraft server:

- `docker-compose.test.yml` (repo root) — boots `itzg/minecraft-server:latest` + the exporter built from `Dockerfile.dev`, wired via a shared `mc-data` volume.
- `.claude/skills/test-minecraft-exporter/SKILL.md` — the runbook. **Read this before reinventing the test setup.** It documents the UID/volume trap (the exporter must run as UID 1000 to read the read-only world mount), the `MC_NAME_SOURCE=offline` label quirk (RCON metrics use names, file metrics use UUIDs in vanilla+offline), and the teardown command.

Do not add ad-hoc `docker run` commands or alternative compose files when this harness already exists — extend it instead.

## Conventions

- Metric names follow Prometheus conventions (`_total` suffix on counters, `_ms` on millisecond gauges); see `CLAUDE.md` for the full taxonomy.
- Don't introduce new flag-parsing libraries; extend `Config` in `pkg/config/config.go` with a kong tag.
- Don't add backwards-compat shims for the kingpin → kong migration. Removed code stays removed.
