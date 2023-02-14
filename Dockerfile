# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ecedc92f903375eb34e3ea398dd488f9b8d08ce81eaf7dec337abac1ba9bcdc3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
