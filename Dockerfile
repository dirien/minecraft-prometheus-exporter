# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ec57fec3e2d529b977ef03539bb80f810222a74494481bec1ffdf5aafbf912b7
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
