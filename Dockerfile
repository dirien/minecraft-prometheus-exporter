# Dockerfile
FROM ghcr.io/distroless/static@sha256:4799cffb309e85b419ffbda31bc272cdce9e87266dcd91769d5c8ba1423916b6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
