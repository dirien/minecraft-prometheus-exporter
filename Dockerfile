# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ae396884b14d3b3181ec128dcfe7c4f6ea428d4cc2923291a51ce1c211e61aaa
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
