# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f62521b319caef384dfa7fa42109549f81a188e2ef8d3b27e351c76334725723
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
