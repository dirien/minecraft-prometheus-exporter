# Dockerfile
FROM cgr.dev/chainguard/static@sha256:fe1d6417c83e3c7cb87456000660720c64ce1cf64eb95fb854c4f92ed324dd4d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
