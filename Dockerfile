# Dockerfile
FROM cgr.dev/chainguard/static@sha256:3ee375b4539717ebcdb96ffc9d94ebb15bc5d950b9d714bd0545997b5aec1c19
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
