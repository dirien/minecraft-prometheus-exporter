# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0fa3935a85aa2349cc89d9715d891c318f700ba951f3945610a2b90c6b0d5e76
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
