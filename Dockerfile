# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9226de08d4ac17a2e14ca529c0585cd411c440c3ff9064ccfffe6e6588ff2651
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
