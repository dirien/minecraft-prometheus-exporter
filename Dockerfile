# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0627bfbabb22c150de61e758b4d0bf52f826e629174056da6433b15232c759c3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
