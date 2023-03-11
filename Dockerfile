# Dockerfile
FROM cgr.dev/chainguard/static@sha256:65adf1ee7fadf72e6f22942841ad0588233982c751fa38dabd8d27a9ff672b4c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
