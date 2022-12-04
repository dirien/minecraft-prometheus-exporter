# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9f92f9dbbdd3a061243ef0915ef431798a0bc0d7a73ee019c110a31cb7f8379e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
