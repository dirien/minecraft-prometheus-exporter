# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2657e6641050ff808452889ae389e3e9f6591d4613156b04862403234d4694c9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
