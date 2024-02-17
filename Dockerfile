# Dockerfile
FROM cgr.dev/chainguard/static@sha256:43f8701510a6cf9cfb1a6805a5970d774309dde2aa5f95bcbb89516f4f4c7a73
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
