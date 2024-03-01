# Dockerfile
FROM cgr.dev/chainguard/static@sha256:67ed8ca8d99e12e8778c038cf88ef7c27d44f08247d317c7135a66ca9d8a7652
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
