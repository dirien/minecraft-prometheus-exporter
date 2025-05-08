# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8d1a96321dca0e8e7848b7db2d431191f15e7e302faa1428100bbab351d42c7a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
