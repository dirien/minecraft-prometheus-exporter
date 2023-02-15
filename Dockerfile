# Dockerfile
FROM cgr.dev/chainguard/static@sha256:cd49355b689f3cde343a9b8c7d3a3cb841070d7147f99144d1ef654905da4594
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
