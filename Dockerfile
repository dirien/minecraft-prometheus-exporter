# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2345a84753de2e9b792e99a2c6ed8cd6a9532e8bc4ceab3e782b27df5801bb3c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
