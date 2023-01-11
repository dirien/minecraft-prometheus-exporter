# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b25ab001448354017969e5abe42496fb107d759d429e69ef69b7da45f0e6cf0e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
