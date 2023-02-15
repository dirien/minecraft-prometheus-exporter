# Dockerfile
FROM cgr.dev/chainguard/static@sha256:53dd6dd6c532ef243d8ae9d56c59d9a46b9c9a801e655cbff3e2abb1acd28596
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
