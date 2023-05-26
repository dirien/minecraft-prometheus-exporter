# Dockerfile
FROM cgr.dev/chainguard/static@sha256:341701ace7231c5dd948d209f5b0d6563ca89cb02699b4e34ae8f93ffa763f38
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
