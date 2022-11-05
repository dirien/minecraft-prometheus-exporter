# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c843b5beee1e85ed6ffb897213ab0833bbcc96bc62ab02f8c84af1aca6fc3af5
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
