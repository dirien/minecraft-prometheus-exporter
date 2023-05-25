# Dockerfile
FROM cgr.dev/chainguard/static@sha256:cd6a89008c4c2190cd0ff15d618e7415ceeed786e2f8c0e44793985133dc4bfe
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
