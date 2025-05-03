# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2a625816afa718bedc374daaed98b7171bb74591f10067f42efb448bfc8ea1ee
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
