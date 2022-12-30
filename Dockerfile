# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4b65af3a0aeed206e931d5c60c555cccdae271194a2ce330db9f9729d3897f07
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
