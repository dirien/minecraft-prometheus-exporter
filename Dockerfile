# Dockerfile
FROM cgr.dev/chainguard/static@sha256:928561b0dbb3c233c9e18a01d37b8bf4789a7402804167f82898ae46427a4245
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
