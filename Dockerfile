# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3469a3c0941fc1ba0b01f07c10cec1c88cfe55e5f001db36777ab8ea2ce4a5fc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
