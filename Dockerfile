# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b4d46e88cac9d0303cd7c42afbe5c4817c9259e637209af15a3492882c6ed840
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
