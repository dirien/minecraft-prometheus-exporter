# Dockerfile
FROM cgr.dev/chainguard/static@sha256:71d6dab9867235fa1e529edd19128021f02c689ab6456ceb26f2b0d8946ff29a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
