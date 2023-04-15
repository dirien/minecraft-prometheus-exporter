# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e0c19add3ef3fae5ff85ff972d0f1462e369bff5e24a1efe71c62659644e5f09
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
