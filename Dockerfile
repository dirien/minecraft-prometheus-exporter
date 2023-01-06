# Dockerfile
FROM cgr.dev/chainguard/static@sha256:720ff01e177f022ddbd3997cc47f942e7433341d7b6b44e9dc76452e3bc658b6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
