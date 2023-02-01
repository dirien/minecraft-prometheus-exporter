# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6e220ebe5a96405019035b476d928b3f8df2ca1cf50d5528d49ab0295c17b85d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
