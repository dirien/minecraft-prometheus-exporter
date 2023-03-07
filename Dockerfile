# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d20729f0549c6e7b1954707a8dec7045a0486daf8c536121b3717afa269b7e41
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
