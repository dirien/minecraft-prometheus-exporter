# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d485a8ed3a3329c5c1c9a3ae99f433528a61abf41d7683776637cdbb1820a2e3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
