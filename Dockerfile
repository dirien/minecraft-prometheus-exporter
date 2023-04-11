# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a020111b8a9725e5ef2a8a90a536bae70e5334b2b668204453b1bd3d9768890d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
