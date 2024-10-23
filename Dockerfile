# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6853717dcd054bc1b641d8a3a31105c0ccbe8d878bba40816d54ad7f1518f869
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
