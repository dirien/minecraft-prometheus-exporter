# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3d302785aef3126a368e4b7d150bc15714cc484d46511f315ca425fdad1f918a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
