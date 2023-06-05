# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bc821fcccb79c77b6f83bd69d07cf7075876bbb0e285d569eed18ffa55fbba12
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
