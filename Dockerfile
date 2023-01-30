# Dockerfile
FROM cgr.dev/chainguard/static@sha256:01dfcc8d455ddb111f496eca6e5fe88aa5efac96a7f08b269f1a259eb5a2f5fc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
