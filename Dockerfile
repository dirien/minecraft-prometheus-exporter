# Dockerfile
FROM cgr.dev/chainguard/static@sha256:289ab2b147f431117275a519a2bcd3b1ceffb54bf386e562596b09cd44be426f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
