# Dockerfile
FROM cgr.dev/chainguard/static@sha256:de1d53e2a83deaf939f8a62095e24ccd867961f54d8bbe67c2f58897be90c53b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
