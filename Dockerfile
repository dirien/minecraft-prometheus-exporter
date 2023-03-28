# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6ba3dbe816174f9845ddc32cd3396a1ebd16be82c1f39cc70b00711de48a9464
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
