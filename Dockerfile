# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c6567cbd31ea3495ef7139e85d239a6b235f65b4393d625fe465ae6ff34cc012
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
