# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c9635595e59e9f4a48da16842ce8dd8984298af3140dcbe5ed2ea4a02156db9c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
