# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e6bd9b0d8c874046fc15093c037f41cd761beaa2d3b0f5961519dd43092578a1
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
