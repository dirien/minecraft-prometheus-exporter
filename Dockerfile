# Dockerfile
FROM cgr.dev/chainguard/static@sha256:cb6724c2348e49e4478a34b9da331498e0f4dd1ebe0cd68462523b6cd73114d4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
