# Dockerfile
FROM ghcr.io/distroless/static@sha256:baf53a1c76184d7e56541a5cdd3f03beea2b22f7112dd71f59dd0b557493290c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
