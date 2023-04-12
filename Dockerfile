# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bd341cccd475faa009a3df7fa7bcbcc4ec59d75fb335df88eb0ba7f21b314038
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
