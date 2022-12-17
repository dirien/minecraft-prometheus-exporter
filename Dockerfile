# Dockerfile
FROM cgr.dev/chainguard/static@sha256:50b8048ffe99f24369e47750f548f3b28264b35e58274156bec97d72649b058c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
