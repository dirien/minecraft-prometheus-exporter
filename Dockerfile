# Dockerfile
FROM cgr.dev/chainguard/static@sha256:24dd7ff8788fdfadda39eeeaefefb6d1cec6002a545935a5f7e017484053734f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
