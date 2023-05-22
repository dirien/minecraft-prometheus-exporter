# Dockerfile
FROM cgr.dev/chainguard/static@sha256:cc63f80218914c469303c5ed8a93e264f1182cb1170b4bbc3396678ba410757e
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
