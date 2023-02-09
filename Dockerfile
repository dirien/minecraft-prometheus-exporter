# Dockerfile
FROM cgr.dev/chainguard/static@sha256:edc3947cea6543079de668b18475cd6335a8c85f191acef5339cba8552859766
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
