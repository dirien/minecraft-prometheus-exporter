# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ca86f743562796423fb33d574811f23a9a83cfbe1369b4d456b61a45114ea195
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
