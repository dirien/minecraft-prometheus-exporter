# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6ee54eb9ffbce221c5ee72f4dab2edad1978dcb023538084dc5fe920b99ae0ca
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
