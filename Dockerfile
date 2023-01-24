# Dockerfile
FROM cgr.dev/chainguard/static@sha256:186682b1fdfafdd5d01f742bae823973119f169bc94f34a195891ee8ee992fef
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
