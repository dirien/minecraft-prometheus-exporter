# Dockerfile
FROM cgr.dev/chainguard/static@sha256:7d8e6efa03a7b58b5a5b2a1d8555e44b990775b29d6324e12d1c77314d595aaa
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
