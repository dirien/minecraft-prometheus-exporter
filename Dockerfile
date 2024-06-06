# Dockerfile
FROM cgr.dev/chainguard/static@sha256:68b8855b2ce85b1c649c0e6c69f93c214f4db75359e4fd07b1df951a4e2b0140
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
