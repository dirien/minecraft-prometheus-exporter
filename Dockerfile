# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8d49dde1a22c75c41b02b38258cebf6a8cc795b0fc7599717b2fdc57652f1388
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
