# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8c5376baf48871a062e2d25e2107b783d518142d9d0c73bcd614d1b67d2d85bb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
