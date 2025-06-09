# Dockerfile
FROM cgr.dev/chainguard/static@sha256:797e62f43d04d792e9f930913e7d9f5a63e92bd19ca5e7e5139a692decee2dbc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
