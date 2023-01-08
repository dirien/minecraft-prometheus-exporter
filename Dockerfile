# Dockerfile
FROM cgr.dev/chainguard/static@sha256:479d52f4d8804ec9cd940e4884c2072bf8a5de58503c0319f2459ad42858636c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
