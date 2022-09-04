# Dockerfile
FROM ghcr.io/distroless/static@sha256:2089052ae16b49a8f38ebc0b201112dcda3fbf6176071d12d2a3690dc67b8f18
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
