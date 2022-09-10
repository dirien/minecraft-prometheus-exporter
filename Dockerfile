# Dockerfile
FROM ghcr.io/distroless/static@sha256:477c99033800e7d30a02e3a293b58eb731f19a48f63d195bf6a06e07a6b00e1a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
