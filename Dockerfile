# Dockerfile
FROM cgr.dev/chainguard/static@sha256:036c737df8f3d79b5b1fa813979d3656651df20ad35d596e074c4c3d2de05707
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
