# Dockerfile
FROM cgr.dev/chainguard/static@sha256:18842dd99c54d96b9ac69f1d0acf73adb2d6478388dcbfa3c759f81a7d0334ab
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
