# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1f3b27eb3810126e06fefad066305669de685de537410c045454aea4eefa3d79
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
