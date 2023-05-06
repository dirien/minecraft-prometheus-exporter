# Dockerfile
FROM cgr.dev/chainguard/static@sha256:af56bfe6ed972fe4c2a901fd33f80c1b43c8974d34516d5b0cc7300e566cb5d2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
