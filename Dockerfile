# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0dccbe0202ab4b4ef046f4f0146afb262f6f3e08992c5ce0635c1e8a4df90794
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
