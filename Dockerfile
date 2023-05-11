# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d1f247050de27feffaedfd47e71c15795a9887d30c76e6d64de9f079765c37a3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
