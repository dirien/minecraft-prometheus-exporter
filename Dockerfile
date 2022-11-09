# Dockerfile
FROM cgr.dev/chainguard/static@sha256:499d1484ef636949e305620d9f4a6f374d1aaef6ff1530bbb1435f3f56953dc3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
