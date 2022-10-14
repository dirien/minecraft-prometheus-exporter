# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a9650a15060275287ebf4530b34020b8d998bd2de9aea00d113c332d8c41eb0b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
