# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e811939c7ab215007bb23b83fb49aee84005b6f6a17aed970a65d9e6bd682e43
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
