# Dockerfile
FROM cgr.dev/chainguard/static@sha256:60214723bfa0298d11ad9f8f149e67f8e556f5b47b7b9cc5557c788941435708
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
