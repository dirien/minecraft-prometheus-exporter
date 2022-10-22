# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8ff5f66d7dba40ed94514e62f947f71805722356f5feb0cf1707addc62c15b7a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
