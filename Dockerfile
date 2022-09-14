# Dockerfile
FROM ghcr.io/distroless/static@sha256:923f5c7f3f9d2f63d01b705464392fc54ecc1172e0c1961e9d725a4dc9b5148a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
