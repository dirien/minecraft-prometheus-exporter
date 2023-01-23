# Dockerfile
FROM cgr.dev/chainguard/static@sha256:7a72a64a71f38521c8ac70b294c54a535a58c88c69c494b6b54903aa06cc76f2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
