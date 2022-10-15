# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9c823791e61f3a4f6e78dc3a1d1c52c150a42a30c64a51bbdf624b65c2a15d0b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
