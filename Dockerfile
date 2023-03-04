# Dockerfile
FROM cgr.dev/chainguard/static@sha256:85fc62be7d63750f68d4790fe10a4ff623a0883a545f050e008d271299f9b662
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
