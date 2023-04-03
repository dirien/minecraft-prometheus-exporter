# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a2c35aeb34f1492f8992ba473451b3a91682fdfd314dbaa8b497fc2760aaa5fa
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
