# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e19d4f4cb05193e9782ccfda43ca30ba7da55b0ca695abdd853ad52b24cb2aa0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
