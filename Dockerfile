# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a126aa44a73d2319d6375ef4cf15a50d9c409db76a87235be6f0c8522afdadd5
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
