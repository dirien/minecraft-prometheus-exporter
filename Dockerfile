# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f5fe67ab41c65f55766d824a2d857a7f56c9058b8e077c43d4d809c467f28df8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
