# Dockerfile
FROM alpine@sha256:28cb8fad9b5faf5bed1a1ecb9fcc4004993c4304bf7a2c2624dacf54caf70902
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]