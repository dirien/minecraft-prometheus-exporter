# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b74f66382b721769876a8dc315b89f14ea9eb914badd3a5145752d6fa2659f0c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
