# Dockerfile
FROM cgr.dev/chainguard/static@sha256:dfcc674e38b2c6816e01cd59a76e1a509ef1156d200eb479f51bbd3e7f84c0b8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
