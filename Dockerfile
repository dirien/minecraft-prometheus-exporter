# Dockerfile
FROM cgr.dev/chainguard/static@sha256:71eec4090c60549046331cc9e52aec92fd8ccc663921041c69655e5943feb41a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
