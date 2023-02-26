# Dockerfile
FROM cgr.dev/chainguard/static@sha256:43c9765db5617f5723c08f5fb8047e6fdfad522b32f090d98bf93b516e9ba5d9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
