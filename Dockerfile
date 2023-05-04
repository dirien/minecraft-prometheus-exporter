# Dockerfile
FROM cgr.dev/chainguard/static@sha256:803916ceb620d392f4bd191a03d23e91cde4400cb55ce404c9b918429b33cc86
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
