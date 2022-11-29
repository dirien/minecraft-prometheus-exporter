# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b74a5e7a36382582394981c3cc6b33e85ca69545936678b4df35ec262c8d5cdd
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
