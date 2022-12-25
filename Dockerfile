# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1fb7a4bccf826fcfbddaa1dc7d5551f1c1214ef55b015ca6a17bf2ad9f43c871
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
