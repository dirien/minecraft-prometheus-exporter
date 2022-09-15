# Dockerfile
FROM ghcr.io/distroless/static@sha256:611e971a974aa75cc7f96bc5ea50c04c5e02bca7cdfa580a133df5352d18b9eb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
