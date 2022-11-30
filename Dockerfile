# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c6cbd5dbfee400a6d19341b6e0bf97f643e0f78d1ec50e5060920f427d2ff4d0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
