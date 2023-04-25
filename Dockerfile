# Dockerfile
FROM cgr.dev/chainguard/static@sha256:261a3c1abf9b71a53c8b3f0029e1b8109be10e66c340839b9a96f9bf3a960026
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
