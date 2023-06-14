# Dockerfile
FROM cgr.dev/chainguard/static@sha256:18bd08608ac704d223f6afc89ed1a7052c44c1b2adf9990384224de94ed76fcb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
