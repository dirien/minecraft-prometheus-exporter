# Dockerfile
FROM cgr.dev/chainguard/static@sha256:df921febce137394c02eaf793c3f00fcbbfc7878b85dd23dfdab1ec067b56e52
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
