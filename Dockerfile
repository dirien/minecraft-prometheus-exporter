# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6b7361ef3fad7bdbf22db570b5d456d3b18c80da1e47539a359387b3b2da1a83
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
