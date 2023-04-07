# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3d0683a03b7173856397403a82d70473d2e98f4fcbca22a56b63a07ef5dfbd16
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
