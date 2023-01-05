# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c8992e967cf1359ca582b3cdbfdcfd6bdf18d7762818ee8ed7843f5743561d25
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
