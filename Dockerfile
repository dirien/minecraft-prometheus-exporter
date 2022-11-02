# Dockerfile
FROM cgr.dev/chainguard/static@sha256:924b952fabe87d4b50a128e55e439db9a44dcad5b008930416bd6c155d518ce3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
