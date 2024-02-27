# Dockerfile
FROM cgr.dev/chainguard/static@sha256:02c5e4a18e90b16c039b4bcfe593bfb0b62373d9e7e9fa53e5de3fb0044d9f77
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
