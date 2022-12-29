# Dockerfile
FROM cgr.dev/chainguard/static@sha256:adfe57c8f44b0645af44938ef88f026cb63fe58506efa7f855d8b1704e91c316
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
