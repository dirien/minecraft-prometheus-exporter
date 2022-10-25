# Dockerfile
FROM cgr.dev/chainguard/static@sha256:867950cc5ef695d5a8771af98b67318e5c3070a2ccdb26d17e4b7a3a1ae5adff
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
