# Dockerfile
FROM cgr.dev/chainguard/static@sha256:caaed002db5b5617771d7fc857bc54f4b0010ef0de98b049a829e9c3b546340a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
