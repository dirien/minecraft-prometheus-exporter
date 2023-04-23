# Dockerfile
FROM cgr.dev/chainguard/static@sha256:364e6b577b58423666f83b8d18ba1a564b7cba54302b9f1708e7887ae69a91a4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
