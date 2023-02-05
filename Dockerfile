# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2ea44d9bdd177a07e6fba8a60f7d45cb8a7358586a5f740187866566e6df310d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
