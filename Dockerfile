# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5ef2713be4309954b594d0b575b746b44b25786ab924894b6b156e73ce48583b
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
