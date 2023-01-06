# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9bf605d88e25a37fbbfe8a54f16e02b8bb070997d5c5100f854a7c29aa90ae56
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
