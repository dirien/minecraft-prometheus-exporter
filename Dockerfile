# Dockerfile
FROM cgr.dev/chainguard/static@sha256:7e1e8a0ed6ebd521f2acfb326a4ae61c2f9b91e5c209dcd0f0e4a5934418a5ec
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
