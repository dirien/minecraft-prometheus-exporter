# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d2a76860057c1260ea5dc8ae4e18beff5ccfb1b67004295c9ab6951833e93de7
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
