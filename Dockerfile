# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2784bc399cdaa1d9fa8c59772fe4ffa96f58dac8784746b291df06b148b59a6d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
