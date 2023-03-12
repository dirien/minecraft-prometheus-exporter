# Dockerfile
FROM cgr.dev/chainguard/static@sha256:221ecadcd681240ade641b133934632c3af3d36e920dac11d977e8d2e92064c2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
