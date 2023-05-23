# Dockerfile
FROM cgr.dev/chainguard/static@sha256:23e1687538e11f8e3edd72a56c5bb4684cb214bbaa6abef90f278ee8860a77f9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
