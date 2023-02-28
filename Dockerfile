# Dockerfile
FROM cgr.dev/chainguard/static@sha256:53527a8a9c2f5acadd8c8c7bbe90234a98106cb4260b6f86f1cb91c61718e7c3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
