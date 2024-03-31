# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8665c8a9fcdab0f8afc09533ee23287c7870de26064d464a10e3baa52f337734
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
