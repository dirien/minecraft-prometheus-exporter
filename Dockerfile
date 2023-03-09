# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c5ae16f127b6a1054b64384a9c95a67a3ffc648bcd44b2e9b6a156b077bb7771
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
