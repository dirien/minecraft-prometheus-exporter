# Dockerfile
FROM cgr.dev/chainguard/static@sha256:110b6918893ea3df0eec04b2f469f3af07e5439900ed259076c55cefb1ec3965
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
