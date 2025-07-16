# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d786d1c686ce4a49376cd2f068d91e691b2bb2e3a6f38513b2396b69b1a9c06f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
