# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ddbb78f5d45b5fc32f5e4a56f98f04a46123c38435e7e9db1989d1852aaecd8e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
