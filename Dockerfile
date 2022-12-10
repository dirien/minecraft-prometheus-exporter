# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f0c505e3a47c5c7593316de0efa5ae21a491198678c2f630dbefae28c78f7d8f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
