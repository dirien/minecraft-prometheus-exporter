# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8999eaa6185ecdce74305a1f530a996df3d368f9b077d3defd08c3638d3e96c2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
