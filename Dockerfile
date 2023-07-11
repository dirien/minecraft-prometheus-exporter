# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6b35c7e7084349b3a71e70219f61ea49b22d663b89b0ea07474e5b44cbc70860
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
