# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9276a4ebe6b98cd1bbd53b8139228434a0e4f00d06d39e33688e9bd759986656
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
