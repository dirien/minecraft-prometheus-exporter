# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a8aeacbaf0a1176ab5dbcf9b73a517665d8db5e1495ba97d64c73b3821deb0d8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
