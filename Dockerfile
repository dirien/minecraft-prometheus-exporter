# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d5d268bbf8d048636123da1c3a3ff97b13e32b46850b8311701c020ccac4265d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
