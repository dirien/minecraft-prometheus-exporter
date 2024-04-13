# Dockerfile
FROM cgr.dev/chainguard/static@sha256:dea7cbb98630ecf732c9d840edec0bda5da5c0c7967a25354fb9f3d8c7f87c1a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
