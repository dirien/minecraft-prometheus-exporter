# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3cc3e55cc12083ef27a6126450319fef45a4380977f3794d39774121fc324a40
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
