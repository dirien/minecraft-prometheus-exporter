# Dockerfile
FROM cgr.dev/chainguard/static@sha256:189b19399010e5b12b67839ad8aec9013db0394ded247d39f7dd4e76160c3d40
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
