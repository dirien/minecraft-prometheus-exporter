# Dockerfile
FROM cgr.dev/chainguard/static@sha256:82fc51cbc8bd3c0d049b0c8b72a8ef02dd633840a2fb9dd0b499725e988087f2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
