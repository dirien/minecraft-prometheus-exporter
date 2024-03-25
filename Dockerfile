# Dockerfile
FROM cgr.dev/chainguard/static@sha256:739aaf25ce9c6ba75c3752d7fde4a94de386c6f44eba27239d2b98e91752e3ff
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
