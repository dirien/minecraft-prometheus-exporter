# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d1c6c919115fa5ba6563a8f641c3b7972856feb2a23a2fe9321df877cf83da18
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
