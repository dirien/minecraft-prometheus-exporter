# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1e14048859430ae84c9e71a56eb53525cb6888e9fc51d262ab633493464ed04b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
