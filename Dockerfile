# Dockerfile
FROM cgr.dev/chainguard/static@sha256:465a411c180be0f864c2c087e209b8c45b7a87a17d96fce431e48dc7b5c18233
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
