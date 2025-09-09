# Dockerfile
FROM cgr.dev/chainguard/static@sha256:7124bf9a6f70e0750d14ef16f1791f322f6d62f50a49223a709f7ed41644c353
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
