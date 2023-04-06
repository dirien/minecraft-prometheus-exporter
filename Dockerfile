# Dockerfile
FROM cgr.dev/chainguard/static@sha256:36987cb01e5cac0a3541e4bf808f46858f5d1c429f2425b63f4decd3cd38e05f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
