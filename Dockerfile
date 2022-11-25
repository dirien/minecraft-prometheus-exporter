# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3224b50f62f316917f8115599741505bb2d0b41b0aa49f3fd39a122d30386e43
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
