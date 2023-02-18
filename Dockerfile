# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9a2320c5820ba0e75cc3a84397dd02b0e36787f24bcdd4a36bb9af5c3a37ec7e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
