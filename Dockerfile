# Dockerfile
FROM cgr.dev/chainguard/static@sha256:999fb8b76c86fc9a2bcfa6c1392e20c549ff0b8cdbbfafaac177c0f038720d37
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
