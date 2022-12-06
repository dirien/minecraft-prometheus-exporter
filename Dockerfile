# Dockerfile
FROM cgr.dev/chainguard/static@sha256:afc2cf1a6ea35a11b27fead4514860c53b1aad41a2140a1308fea4debe97b0a0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
