# Dockerfile
FROM cgr.dev/chainguard/static@sha256:51558898680cc546996e628e26fa6b550ad10e31074420e6de7e605e9127a861
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
