# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ec82f35b2ab07d8056b652032f9c8386951d8b4a41ed42722d1cdefe0a6e1086
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
