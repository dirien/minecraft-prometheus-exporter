# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d30d21c0cdb76d4af27f3495597c6e2a6ddc78f8addba170785d61f548dbdf17
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
