# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f3c59f6f175ee9a9f4fec46a24c86a41b10abdc900a74f823e0c9ece19089afc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
