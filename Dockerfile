# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d835f471249bb573085015d7812828cea98643164649c8ab7438a55beac2a237
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
