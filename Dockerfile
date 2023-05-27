# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ee47224a2afc674c1f1089b9dea97d5ee400cf2fff3797398778450a4cfb2a8d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
