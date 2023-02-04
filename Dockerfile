# Dockerfile
FROM cgr.dev/chainguard/static@sha256:98785c6b139ce7023264a7f42d4658624c53e3ec361538d2b98547da092b56e0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
