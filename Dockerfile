# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5497b01f36ef14a5198c0165e50ae6a0006d0c7457d4566f1110257e1c0812ed
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
