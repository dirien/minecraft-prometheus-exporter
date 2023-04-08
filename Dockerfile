# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d6a85300ea1a7e58767249795b3feb8e60a33d60dfa3906c57a7424ffca45dab
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
