# Dockerfile
FROM cgr.dev/chainguard/static@sha256:668e8e0b02e783df036c3bd115b088b5e27e4647829475e4b133797202814b5a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
