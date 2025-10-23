# Dockerfile
FROM cgr.dev/chainguard/static@sha256:01f45a2a6b87a54e242361c217335b4e792b09b92cd4b0780f8b253e27d299bb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
