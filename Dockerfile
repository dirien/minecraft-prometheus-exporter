# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0fc4584e8ae2a6e71f0bf73213c659f64756b48caab1bba3207c2ea51d292fb0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
