# Dockerfile
FROM cgr.dev/chainguard/static@sha256:008fba246fcf09e99d764f1bff5ceb4aeec2be476a5a3d40c56c991274d70a40
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
