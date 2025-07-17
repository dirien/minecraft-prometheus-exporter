# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e55a04f85f58e6b0e36bae05b8ff18c79035e65c0151e4e866eb49679782d28e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
