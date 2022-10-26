# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f4f220edda7865cba10110566c829a45e5051b405f4c085a2db24a86d4ac7525
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
