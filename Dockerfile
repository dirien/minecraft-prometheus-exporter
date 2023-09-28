# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ef5add7fd46cf1ce7d33d6de517833ac5c7e749db9b15249f9c472a772f3af27
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
