# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a117fb6e8c62246fe60e40eaa0c2cb51575db8cec42c114bc5a9d4cb89d94fee
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
