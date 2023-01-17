# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6b5fe3a1806308e5eecb8e6b23404b3865099dde9c6542930d6f0f8776a54499
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
