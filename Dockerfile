# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4da027e241598d68296061cd9a20ab835a8520bfd969e7f71352cb27a6a3d08a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
