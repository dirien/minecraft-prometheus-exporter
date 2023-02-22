# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4fea7521cc3dcf49cb6733f4cec86237f4396558dfb9d16b91fe71f727ba31e6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
