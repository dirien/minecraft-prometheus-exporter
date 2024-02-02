# Dockerfile
FROM cgr.dev/chainguard/static@sha256:fdff80538f4720a302de41af54d1491060b6366070622aad41b7253e8a9d2879
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
