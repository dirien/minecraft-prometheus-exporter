# Dockerfile
FROM ghcr.io/distroless/static@sha256:1972ba2d5e5d917438f1a4d12cf832358ecf9d5df5cd45fe106d0afae6d8a3c0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
