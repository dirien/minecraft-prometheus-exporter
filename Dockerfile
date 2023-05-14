# Dockerfile
FROM cgr.dev/chainguard/static@sha256:fa2dd5071bfc476e04785aa83472f0b67e098b6baad5bf21249beb250abb56a6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
