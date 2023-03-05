# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4a5fda9b2aa55b49971d220cc4ba3d73998084e37e437f23721836112015c2d4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
