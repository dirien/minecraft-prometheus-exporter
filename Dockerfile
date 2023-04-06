# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bdefc3bd603d6a71a8e18812cdc6f69cbb6a21e7f165010a6c2e512e2f7ad8f6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
