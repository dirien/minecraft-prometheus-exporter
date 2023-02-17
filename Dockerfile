# Dockerfile
FROM cgr.dev/chainguard/static@sha256:15a59e3da3b97288eaecce6c9ac00df05a9f4c261bfbf56a38cfd688f5000f95
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
