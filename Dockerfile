# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f105d07cd4f3d9cf4764c5812c6053ca7948eaf595adab67c20a8aa28349434b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
