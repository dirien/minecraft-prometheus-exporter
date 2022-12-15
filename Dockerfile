# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e8e2bfce2c44d413ca5f1346e3e4faa0e28b01a1afe8d16169ce4fc552193a44
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
