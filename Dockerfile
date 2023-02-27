# Dockerfile
FROM cgr.dev/chainguard/static@sha256:eafb37684f58989f96709f2dbc4346580f0fc867ff164c2252a2d280e8096e7f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
