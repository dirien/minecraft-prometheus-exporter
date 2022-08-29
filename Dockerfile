# Dockerfile
FROM ghcr.io/distroless/static@sha256:88c03d1613bc2fe846a1554f41beac3bfd8479a9853a29ac55b73088333990f8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
