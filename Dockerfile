# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1856b37dea22fbb62cd5bb00280b9858d5d5726d4a5a923d992ad389fdd95474
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
