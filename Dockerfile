# Dockerfile
FROM cgr.dev/chainguard/static@sha256:fce1fdce595332afe7f339303ae288c349c4e4139d926f8c7ccf4b7ca2911553
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
