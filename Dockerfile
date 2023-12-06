# Dockerfile
FROM cgr.dev/chainguard/static@sha256:ab062ebcd496faecdec3961b0e8061d81ce1553595432a7e6d212ff2c3bd46d8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
