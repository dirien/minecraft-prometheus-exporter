# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9beb8782060e6c4737b9c41ecc82bab4f81ab1fa7d38269b5f93e06b6622c3b0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
