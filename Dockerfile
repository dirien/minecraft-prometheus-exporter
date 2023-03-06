# Dockerfile
FROM cgr.dev/chainguard/static@sha256:508cb9c9888647f74cc16bb255563069e7ffbf7cc753ef209fe46118b6a2359f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
