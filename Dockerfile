# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f4df02404beeb36e8a8440339818f2dc3d577dcef807a3f74df631e6dbf3a661
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
