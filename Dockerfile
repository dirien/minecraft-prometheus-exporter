# Dockerfile
FROM cgr.dev/chainguard/static@sha256:859bf39e084260b8d6f52dfa7a5a4892e2fd0b580d25d9fc19de982eb85c9333
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
