# Dockerfile
FROM cgr.dev/chainguard/static@sha256:483f841e0b8767a48dac0c3317f3f89eaa40a54da4f525e7c2e8eeadfebc3e6f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
