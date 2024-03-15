# Dockerfile
FROM cgr.dev/chainguard/static@sha256:45346401d4619bb419d412e9f6cf6723e2a914219c4adc21c699d7616cfe9c69
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
