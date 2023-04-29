# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ed77c0cde426c6d48aa27040d5217ee4888d5d441eeece0d5ad66279d6fe65f7
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
