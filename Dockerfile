# Dockerfile
FROM ghcr.io/distroless/static@sha256:5604d1c37a1a59efbc16163426af8d6a33524625769230ee95f204ea6957daa8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
