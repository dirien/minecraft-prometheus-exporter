# Dockerfile
FROM cgr.dev/chainguard/static@sha256:696ab4052bac35c7d1c8ef0f2be55f54b05ce959d14327b8ab1af14f5da01440
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
