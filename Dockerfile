# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9ba2ca5edb120a25e6055e9f88544314991a1eb52ae1390b9d372e2570c8bc9c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
