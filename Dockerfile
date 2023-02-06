# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5adf751139e27adcd9d8519206b383c37ff0525dab0cf33d466651ae2d8e9e8f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
