# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4a58d5bcd0dcae2bfebf28f7ed9621c6d21838a5f343cdb9ee9450f1af478591
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
