# Dockerfile
FROM cgr.dev/chainguard/static@sha256:beb6a9eaf915a03a6fedbeda117fd327cd6b08883ae5fa58bd2ac7c0980318cd
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
