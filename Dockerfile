# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1e81a8138b8321c39d5cf4cbc32b270f3852f1fb133f9865859d7ff683609ce4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
