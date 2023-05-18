# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f169ffc4acfad23d529a40d74c3424a22e55417c59a330cf98e4cfdd62201c87
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
