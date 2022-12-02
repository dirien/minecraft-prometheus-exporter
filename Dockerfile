# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0c323a92c7586a59a380d90053bf960db6a374c12388eb13035f9ba3c0871830
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
