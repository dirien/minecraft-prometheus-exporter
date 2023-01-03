# Dockerfile
FROM cgr.dev/chainguard/static@sha256:961b4e06b6578dfa943333cd9fcb15d9cd64562d54e69e3de0f0c1f1a2ba2fc5
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
