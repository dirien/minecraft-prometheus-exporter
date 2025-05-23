# Dockerfile
FROM cgr.dev/chainguard/static@sha256:48278935856fba0e9fac80365ae9a5b33297f7e5682c2dcb86ecfe5eb6878972
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
