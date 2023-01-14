# Dockerfile
FROM cgr.dev/chainguard/static@sha256:51f4132dac6204540f004d5f1835fc51426a154312073d8dcc5746f8381bc9d9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
