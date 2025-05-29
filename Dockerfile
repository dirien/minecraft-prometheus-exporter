# Dockerfile
FROM cgr.dev/chainguard/static@sha256:633aabd19a2d1b9d4ccc1f4b704eb5e9d34ce6ad231a4f5b7f7a3af1307fdba8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
