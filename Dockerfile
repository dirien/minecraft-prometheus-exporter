# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5ba1fe7f2fbf7467bcfe0caea261502ee7ffb503b6524481128b22584f2131c0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
