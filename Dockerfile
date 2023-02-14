# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f35579d03a5a2c54752dd21430f20172455ba0514327a41a135501f067ce0358
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
