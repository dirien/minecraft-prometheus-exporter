# Dockerfile
FROM cgr.dev/chainguard/static@sha256:aa1aefcd33abda25790de485578ae302c82690fd037d62876f85586e1af4c4d3
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
