# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b3af1cd6be9a6c16497b4b9e0df73fa5a8a3a60bd1831433f8c0bfa879e62b97
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
