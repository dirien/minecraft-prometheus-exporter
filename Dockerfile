# Dockerfile
FROM cgr.dev/chainguard/static@sha256:927b01352666b5e72a71029ff78bec8a35c069abf8733e6ec509318f942e6527
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
