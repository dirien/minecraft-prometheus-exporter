# Dockerfile
FROM cgr.dev/chainguard/static@sha256:afaa9c3ad3772105e115fe7fff5e8ef9909e7592777822930121df68c57d08cb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
