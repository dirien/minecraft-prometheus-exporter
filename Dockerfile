# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5f4fda06f9daa9c7f421c6652fb7b7e5392feb1f07e52355e4b5057e3ea60daa
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
