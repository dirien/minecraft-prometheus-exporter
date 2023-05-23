# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6100d7cc1dc23ba7447393345838f3b3ece188adf6fc0160409cb70a80f33981
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
