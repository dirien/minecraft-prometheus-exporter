# Dockerfile
FROM cgr.dev/chainguard/static@sha256:96681f4eb60a0d6ba65703a610c0f27a6935073ea49d2c2585ddb93815a90b23
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
