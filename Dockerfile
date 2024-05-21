# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ad8d1dbfff508ee09c6aaf1c5227717ef3dd3180769c8e0ba90bc1fcf553258d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
