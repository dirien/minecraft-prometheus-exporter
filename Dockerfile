# Dockerfile
FROM cgr.dev/chainguard/static@sha256:17844d8faa68296eddddad9c1678edaf18cc556433d68c5bd3808a6efac200a8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
