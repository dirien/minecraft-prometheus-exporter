# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9047a7c9e2d7dbc6614053e82c859e0c409591ccf054447fdc375a5e5741f70b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
