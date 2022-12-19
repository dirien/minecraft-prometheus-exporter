# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ab7368ad9afbc3bf3ee190f20833c4e88cc09dc1ff3ea451620f3d0d55d2c4d8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
