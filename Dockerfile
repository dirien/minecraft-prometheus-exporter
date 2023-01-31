# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6d35065aefab8259638ac69194fc20561f1c2675a0d64f21f3715c4639da4207
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
