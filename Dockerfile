# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6a4b683f4708f1f167ba218e31fcac0b7515d94c33c3acf223c36d5c6acd3783
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
