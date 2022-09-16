# Dockerfile
FROM ghcr.io/distroless/static@sha256:d3cb3af96f29265cacaabdc61951bf5f04036e1e44084b2d1d6b270a2969388e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
