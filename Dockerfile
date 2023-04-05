# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5cf9767b6dde8f647903d5fdb19c343cbc941a0c64f333f42916ef4f22ed6702
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
