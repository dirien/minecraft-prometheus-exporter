# Dockerfile
FROM cgr.dev/chainguard/static@sha256:70b9e913abf2534b7f80a958261bc807c2f5bc31c4f9ba8f901046c70c83b2ae
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
