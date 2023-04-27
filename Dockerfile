# Dockerfile
FROM cgr.dev/chainguard/static@sha256:505f1e86fd38d5936d39e78f05e15462c5030afa94022f9c5078bfe2c64bcf59
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
