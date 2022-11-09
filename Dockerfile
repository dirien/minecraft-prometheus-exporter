# Dockerfile
FROM cgr.dev/chainguard/static@sha256:15e433b8374c1bfbe3049fca2cec1ba828e185159dda8884f999c146c09f8698
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
