# Dockerfile
FROM cgr.dev/chainguard/static@sha256:2314ce91dbef3a9a1100a0992b16c4ca773b793586710d9c5c4cffb617f1178c
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
