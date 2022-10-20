# Dockerfile
FROM cgr.dev/chainguard/static@sha256:7b37e6e1a5fa6f4e246b33b5797e2c90a34838fc3b3aff4ee7c627a5a583ff28
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
