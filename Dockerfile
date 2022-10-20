# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e2b30b69df819debb34e6f99e4c533c113883a5dec2af61d6fb42a6ebc81b783
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
