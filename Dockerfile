# Dockerfile
FROM ghcr.io/distroless/static@sha256:51de142ee459a23e37377eedd23743bfd6a606ffca0e48aac8f599db7bfa72b1
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
