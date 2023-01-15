# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2a8599a3ae6a2cfc31056b98552272df8f50e9df1780490c958a6a4bd79e1d1c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
