# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d88e69947068712bd269626bca96bf250d21581b7066a67da746813ba8260a9e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
