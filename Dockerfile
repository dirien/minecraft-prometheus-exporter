# Dockerfile
FROM cgr.dev/chainguard/static@sha256:81018475098138883b80dcc9c1242eb02b53465297724b18e88591a752d2a49c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
