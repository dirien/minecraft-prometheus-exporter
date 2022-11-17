# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f73591546d523b50d64596923d919d0dad04e1fc8e23c0fb30f53a0b73d3ea05
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
