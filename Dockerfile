# Dockerfile
FROM cgr.dev/chainguard/static@sha256:326f91770528ede0c3e3413f7dca5a278f637a0c49aa32ce1a68a75cc67afdae
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
