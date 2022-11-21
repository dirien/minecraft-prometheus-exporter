# Dockerfile
FROM cgr.dev/chainguard/static@sha256:26f7057ac20af5735200b948533f45032026cd9a1be0afa8b344a5761d303ad7
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
