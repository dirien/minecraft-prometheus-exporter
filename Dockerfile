# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2fd46e35dfc0c0fa72bdbe2526dd3b80dd9d05b3b1a3df87e4c5defe414e3273
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
