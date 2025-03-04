# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3e9af2550ae5ff1fe5b9d69332955c01213c37c75874b184e5fbea500d1c9808
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
