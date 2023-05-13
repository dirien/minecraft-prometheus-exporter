# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1ab0a290ac6ef0ad24cd179e39d2e8351d09eeea22ea0013998f525b7fd8d872
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
