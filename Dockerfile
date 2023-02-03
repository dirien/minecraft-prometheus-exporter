# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a218b8525e4db35a0ce8fb5b13e2a980cc3ceef78b6bf88aabbb700373c1c2e2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
