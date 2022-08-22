# Dockerfile
FROM ghcr.io/distroless/static@sha256:ca249938b61009fa5c08b5a54d7103a20b2a376dd8965b694e5293c146317e43
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
