# Dockerfile
FROM ghcr.io/distroless/static@sha256:4f0407a6e4343b6b5550742a246eecfead4c2b0c39ec0bb2553e630fe5398372
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
