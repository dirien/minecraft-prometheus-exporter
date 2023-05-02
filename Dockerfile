# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d5138c533347d722b3e923677d66d6c34844756371fb3e82f64b745ae23d7eeb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
