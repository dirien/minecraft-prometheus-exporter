# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ff121d5cb08232edd4896cdb2d0e56349d2428909e194514ad3c0e1dd57874d9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
