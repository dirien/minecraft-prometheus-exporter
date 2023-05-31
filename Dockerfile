# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f7914e3925ee3ff1355f713ed931a01d8036767a08a6a80af60964c249a65dff
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
