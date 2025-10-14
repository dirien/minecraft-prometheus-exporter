# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b00a88ca2a8136cdbd86a8aa834c3f69c17debb295a38055f7babfc2c9f9a02b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
