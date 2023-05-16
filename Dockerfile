# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1888f4db2c92e5a3e1b81952d8727e63c1b5b87ad3df374de318999beb4fd194
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
