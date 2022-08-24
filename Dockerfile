# Dockerfile
FROM ghcr.io/distroless/static@sha256:5120ede7ee241d9a7390fe01dd1f4c43ee782a22073df80fd7d9e620b78fe7e9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
