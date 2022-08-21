# Dockerfile
FROM ghcr.io/distroless/alpine-base@sha256:73d3cd83c1a54d6b3445e9e97dbad39b345a8b91319f39c0447951d1022822c3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
