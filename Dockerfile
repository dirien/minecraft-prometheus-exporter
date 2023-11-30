# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e0f163b457130b20176fb5ffd3a5528ee69c314ed02a0e9b5a4047806214f832
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
