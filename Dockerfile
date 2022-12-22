# Dockerfile
FROM cgr.dev/chainguard/static@sha256:644dbd4965dd00bfc7df110db218437bc8907d3b37fca281ea8844ba240b759b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
