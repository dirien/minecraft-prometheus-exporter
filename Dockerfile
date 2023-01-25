# Dockerfile
FROM cgr.dev/chainguard/static@sha256:441f1469b9a05ca1954a00368513dcf319695b098d5dbd9a7d9d6ca0e87bf240
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
