# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d6621cc6a8d7d6504c701dce729b020b74e696af1715335b3390dc194f2195c0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
