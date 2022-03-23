# Dockerfile
FROM alpine@sha256:6af1b11bbb17f4c311e269db6530e4da2738262af5fd9064ccdf109b765860fb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]