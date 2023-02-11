# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1b111fc8c261a56f85606688fd9f182411925cc6ecf9186e59520cdbb65acfd6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
