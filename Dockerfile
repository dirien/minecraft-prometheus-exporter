# Dockerfile
FROM ghcr.io/distroless/static@sha256:612d1196d8d3a690e01b0846d5a2f48ab16e4b3a62d3bbe7d9e8743deb143049
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
