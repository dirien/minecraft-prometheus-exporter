# Dockerfile
FROM cgr.dev/chainguard/static@sha256:692b379713e1987102faac71b996db54dece252e10c28f62b15fe35b7f2e8257
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
