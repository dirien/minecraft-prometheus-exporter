# Dockerfile
FROM cgr.dev/chainguard/static@sha256:54b589146d4dbc80a094fcbcd6b09414f3df94cde8ea6d31c44fd02692c58203
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
