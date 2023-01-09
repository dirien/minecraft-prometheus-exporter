# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c2531047f47e0964cb4f9ce93c9ad65dcea9926af42070d870b7a301b14256db
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
