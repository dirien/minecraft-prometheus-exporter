# Dockerfile
FROM ghcr.io/distroless/static@sha256:e577a44fd72b647a34d5d1c4e0686c2f2a6837d4ee985f29b073b03d7e5903a1
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
