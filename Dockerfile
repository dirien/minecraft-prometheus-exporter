# Dockerfile
FROM cgr.dev/chainguard/static@sha256:873e9709e2a83acc995ff24e71c100480f9c0368e0d86eaee9c3c7cb8fb5f0e0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
