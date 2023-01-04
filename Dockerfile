# Dockerfile
FROM cgr.dev/chainguard/static@sha256:be400c5f4fa4d8b07fa6a4c6f0e264403f348c7b13606d49fef3f5b5e97f16d8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
