# Dockerfile
FROM cgr.dev/chainguard/static@sha256:890c498d6a073e31ff41f136fcf5266cc13a85cd2d2264b17a200e1d9880bc57
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
