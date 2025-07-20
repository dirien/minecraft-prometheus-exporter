# Dockerfile
FROM cgr.dev/chainguard/static@sha256:93b70336be10c325d5a96016971b71b38d8e79e5148af2144f2aae93ee9367c3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
