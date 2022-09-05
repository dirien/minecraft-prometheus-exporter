# Dockerfile
FROM ghcr.io/distroless/static@sha256:4eadd48e6984614cc5be91ee3d39e1c59b1d74233e88bf4d33118e343cab47f7
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
