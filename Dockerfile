# Dockerfile
FROM cgr.dev/chainguard/static@sha256:05b3f5769206e07d8c70c7314eda3fab4cc084bd5d0576ecc6a1c5b441fc35ea
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
