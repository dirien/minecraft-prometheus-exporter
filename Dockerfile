# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2931a5b1e80439b89f173615ce2d3286027e0b617fce48a0ae77e7ac3c93f221
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
