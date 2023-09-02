# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
