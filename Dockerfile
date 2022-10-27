# Dockerfile
FROM cgr.dev/chainguard/static@sha256:cc197ed23a6203f7dc10fb0cf19ed44488c3e44fce2ff3885c8d546d3c97d760
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
