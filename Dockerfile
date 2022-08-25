# Dockerfile
FROM ghcr.io/distroless/static@sha256:411ad5c3acd92fc3625c3e2e3540f85f4773123e6dae07bb3a9b330bd447c51a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
