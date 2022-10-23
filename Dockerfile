# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6b57ba856e5cdf200e9028aedfb48d850f911e4dc55249b3060a043202660130
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
