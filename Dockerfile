# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d624beaae60b10ec342896b1470421eb473549cb3ac3bf1c6ec074d8b460b4fc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
