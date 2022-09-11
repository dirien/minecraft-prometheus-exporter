# Dockerfile
FROM ghcr.io/distroless/static@sha256:1a4ae9b4d7fb982586ad329b6a416b4d5201871d3a2e56d117c96ae2129aefbe
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
