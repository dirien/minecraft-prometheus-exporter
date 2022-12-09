# Dockerfile
FROM cgr.dev/chainguard/static@sha256:7e74b3d8813aaf07dd7d2961dc1748928ec9f89e301a0ff23253023370938e90
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
