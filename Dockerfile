# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e23dcac1afb48897e466f95f42a293ae066ee5f7d2c27c2ede5259a3b590ef6e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
