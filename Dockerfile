# Dockerfile
FROM cgr.dev/chainguard/static@sha256:aad026c3b37cbd3b26097ffc641e27c7d99220a8a25f07f18f8a190051652f10
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
