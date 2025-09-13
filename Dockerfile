# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b2e1c3d3627093e54f6805823e73edd17ab93d6c7202e672988080c863e0412b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
