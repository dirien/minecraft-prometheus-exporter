# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4780ca5afa898980cfda7029ab3ab6fd958764bc1e48d57a2713514c1e68887b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
