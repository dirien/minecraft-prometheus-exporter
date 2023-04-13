# Dockerfile
FROM cgr.dev/chainguard/static@sha256:9cbecd5f1e1b0ee97b7bfae2d110d0b0fb7fbba9a5198ce1f6aac46e544696c4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
