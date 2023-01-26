# Dockerfile
FROM cgr.dev/chainguard/static@sha256:052a77bdcff373d829fb504c495782852946cb20f9f72385e771670d61b291f3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
