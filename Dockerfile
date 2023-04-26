# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6eba22e915f5859b03512dba29eb3ff6ae63a940c83783a0ac64eb0fb58cc356
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
