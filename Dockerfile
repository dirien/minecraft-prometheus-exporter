# Dockerfile
FROM cgr.dev/chainguard/static@sha256:303006d4ce93d8038be20728389f228436f9ec34450e857baab42ec91bde56cb
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
