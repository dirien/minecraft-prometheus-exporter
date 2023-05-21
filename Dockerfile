# Dockerfile
FROM cgr.dev/chainguard/static@sha256:99bc5f9865a0e2c4d0b4ec94e284cccac4c73d5b50d660be327878dc3f3e8df3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
