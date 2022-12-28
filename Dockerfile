# Dockerfile
FROM cgr.dev/chainguard/static@sha256:977285fd4bc0d8792975371511fd7ff1ebad97a24c21c1b9dba15c680af6a16e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
