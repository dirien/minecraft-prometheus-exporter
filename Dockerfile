# Dockerfile
FROM cgr.dev/chainguard/static@sha256:aface1c7cc4f2ef13efb3d8579c9acbdd0d8c5670f2daab98fd1199c4cd03538
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
