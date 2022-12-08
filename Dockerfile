# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f89f7e7b28a8a4b8cf467e9712d9675b54ae9571329f63db4e96f785b666914a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
