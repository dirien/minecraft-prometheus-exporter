# Dockerfile
FROM cgr.dev/chainguard/static@sha256:52b025c8c15c779d3ba93a9dfead8000e302d288685d735ce7daa4b712ddf46c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
