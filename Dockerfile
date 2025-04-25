# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2e3db1641bb4fe4e85d2210f4aadb79252e90d5fa745f53a3ffed6a1aab4f73b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
