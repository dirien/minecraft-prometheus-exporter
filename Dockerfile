# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a2f525dac2f9ec900283ead64eb88a6037b2989630615ee8de8a2dc7bfcf152b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
