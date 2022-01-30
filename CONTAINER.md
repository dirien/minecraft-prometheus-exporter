# Minecraft Exporter for Prometheus

![minecraft-exporter](https://dirien.github.io/minecraft-prometheus-exporter/img/minecraft-exporter.jpg)

![Docker](https://img.shields.io/badge/docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Alpine Linux 3.15.0](https://img.shields.io/badge/alpine_linux_3.15.0-0D597F?style=for-the-badge&logo=alpine-linux&logoColor=white)

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/minecraft-exporter-image&style=for-the-badge)](https://artifacthub.io/packages/search?repo=minecraft-exporter-image)

This is the documentation for using the Minecraft Exporter for Prometheus as container.

### Installation

```bash
docker pull ghcr.io/dirien/minecraft-exporter

or

nerdctl pull ghcr.io/dirien/minecraft-exporter
```

### Running

The container is running on port 9150 by default, so you need to bind the port to your host.

```bash
docker container run -d -p 9150:9150 ghcr.io/dirien/minecraft-exporter

or

nerdctl container run -d -p 9150:9150 ghcr.io/dirien/minecraft-exporter
```

You need to mount the location of your world folder to the container. The default location in the exporter
is `/minecraft/world`.

```bash
docker container run -p 9150:9150 -v /folder/world:/minecraft/world ghcr.io/dirien/minecraft-exporter

or

nerdctl container run -d -p 9150:9150 -v /folder/world:/minecraft/world ghcr.io/dirien/minecraft-exporter
```

### Additional environment variables

You could use the flags too, but I would not recommend it. I prefer to use the environment variables.

| ENV variable                  | Description                                                |
|-------------------------------|------------------------------------------------------------|
| WEB_DISABLED_EXPORTER_METRICS | Disabling collection of exporter metrics (like go_*)       | 
| MC_CONFIG_PATH                | Path to YAML file with config for the mc variables         |
| MC_WORLD                      | Path to the world folder.                                  |
| MC_RCON_ADDRESS               | Address for the Minecraft RCON.                            |
| MC_RCON_PASSWORD              | Password for the Minecraft RCON.                           |
| MC_NAME_SOURCE                | How to retrieve names of players: offline, bukkit, mojang. |
| MC_MOD_SERVER_STATS           | Additional server stats for papermc or forge               |

### Legal Disclaimer 👮

This project is not affiliated with Mojang Studios, XBox Game Studios, Double Eleven or the Minecraft brand.

"Minecraft" is a trademark of Mojang Synergies AB.

Other trademarks referenced herein are property of their respective owners.