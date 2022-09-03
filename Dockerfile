# Dockerfile
FROM ghcr.io/distroless/static@sha256:dce7e482f7d8a27b0003569ccbc0b42fc5167deac7b56f505035f557aee165cd
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
