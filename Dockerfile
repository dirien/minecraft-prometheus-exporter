# Dockerfile
FROM ghcr.io/distroless/static@sha256:1a83c8dae2f71e0c95d302a98a1b12a732ccea60c8f7feb226fa70615ca001b6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
