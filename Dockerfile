# Dockerfile
FROM cgr.dev/chainguard/static@sha256:dd37b804e4de19f7d979c719e3520a79651d0158cafc27fe2005d099a0029afd
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
