# Dockerfile
FROM cgr.dev/chainguard/static@sha256:97a544db65bdc6e71fd9e2729e666affccefdff0e223d30341e898bc22cbe0f6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
