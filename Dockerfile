# Dockerfile
FROM cgr.dev/chainguard/static@sha256:39ae0654d64cb72003216f6148e581e6d7cf239ac32325867af46666e31739d2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
