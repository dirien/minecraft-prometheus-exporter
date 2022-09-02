# Dockerfile
FROM ghcr.io/distroless/static@sha256:3003bbfa313fbaab06980a636dc19e5fc1e1818b60bb1924dcfc69b134d7e4a6
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
