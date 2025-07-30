# Dockerfile
FROM cgr.dev/chainguard/static@sha256:81b61e16687f76ebc3c1fa71ec3fa3e0901e2908e0cd442f378557c294920aac
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
