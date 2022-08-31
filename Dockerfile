# Dockerfile
FROM ghcr.io/distroless/static@sha256:ee42ceccc576840924d31396c229575596f3cda32d4faf2dfd0abd72ec6a8dec
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
