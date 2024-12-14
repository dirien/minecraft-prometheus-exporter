# Dockerfile
FROM cgr.dev/chainguard/static@sha256:211178c05cff513767ad12d2bcce7e641e2772865ddbea0e7394afa12725d50a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
