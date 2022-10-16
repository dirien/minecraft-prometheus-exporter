# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f6110e38c195a5f2d53e72533c7ede9c46c262ada58f945e97b362381f410da8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
