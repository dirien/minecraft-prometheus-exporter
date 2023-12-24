# Dockerfile
FROM cgr.dev/chainguard/static@sha256:296c640b19594bafcce3bb20ccf331b50998dc536dd9824c4684d819b7364615
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
