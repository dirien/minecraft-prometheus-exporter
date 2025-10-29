# Dockerfile
FROM cgr.dev/chainguard/static@sha256:939a132511fcbc2702e0e251b6f3ea368c0ad4f114678ae5973903352357d01a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
