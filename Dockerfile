# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1f1087961ccd9fde388bdc45a3c38952a34ac2b333421afbefa2f4da771bb8b8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
