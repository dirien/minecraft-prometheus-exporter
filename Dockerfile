# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5b72996ac2b7d7c6e26ab38c6ffd172b9d85f14cd1f0cb852a2a4f4fbf05c82e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
