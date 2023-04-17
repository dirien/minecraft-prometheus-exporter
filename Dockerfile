# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bcf3360cf1a5119803e049ccdd6d420aad713163900c4d2c6dbdea3e28da9df6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
