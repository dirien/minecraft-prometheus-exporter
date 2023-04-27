# Dockerfile
FROM cgr.dev/chainguard/static@sha256:53057e0dcc945d74c9126e04d5d128edca13c19991640e632d61f18fa7afc659
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
