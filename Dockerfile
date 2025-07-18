# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3c0cfe403ce6f1ff7761aab482448e5d4979762a9853103639c0d769aa7de89e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
