# Dockerfile
FROM alpine:3.15.0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]