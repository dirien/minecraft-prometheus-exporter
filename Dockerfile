# Dockerfile
FROM ghcr.io/distroless/static@sha256:a0b4cb702c37ca2a667ea98356620a5dd0ca51c5ec459b34331f49f842b00236
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
