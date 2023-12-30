# Dockerfile
FROM cgr.dev/chainguard/static@sha256:37200bd74dac036a7a40f1984eefa2123d0705540daaeda6b1faee727db66726
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
