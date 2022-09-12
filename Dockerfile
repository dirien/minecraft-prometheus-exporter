# Dockerfile
FROM ghcr.io/distroless/static@sha256:a65803353d3a91249b1698cdacd6cf4169678a3654060050530ccf621102b877
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
