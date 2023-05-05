# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b60e384c5182f2d0818b1d71cccf921f5eb9d72cfab0dc2aebd8f26311e2e4c1
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
