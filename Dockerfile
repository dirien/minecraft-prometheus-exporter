# Dockerfile
FROM cgr.dev/chainguard/static@sha256:01d6e4b626cc452b85cfafe427b8ea8cffe963fc5cccbae5684a5d81542d753a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
