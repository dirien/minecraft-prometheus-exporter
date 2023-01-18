# Dockerfile
FROM cgr.dev/chainguard/static@sha256:07b84aa60aeae795844e89c3c53cc3309dfcb5c6c69e3cc8f29564187b8fa497
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
