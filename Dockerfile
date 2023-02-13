# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1907d49419f023f7843fd3592f940d1564148c3cead10e7c814e6dea61c0ce9d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
