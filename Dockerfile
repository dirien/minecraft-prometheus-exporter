# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3cdfd83ff1454f633b3ad76927ba896f7adb921e53e57d3b82bd2b2d6f558a1e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
