# Dockerfile
FROM cgr.dev/chainguard/static@sha256:23dafa56657d33aa447aaea99c20b58c77ed193e8a1b84ed43a3ff12ab79b080
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
