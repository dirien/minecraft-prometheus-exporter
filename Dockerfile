# Dockerfile
FROM cgr.dev/chainguard/static@sha256:87c802c6011260141ecd54425a2ee498f4f8260ad77ce423a364aaf1a5840a27
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
