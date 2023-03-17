# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d4bab5696fb2cec9d1cf184b573ef2bf098521e98168364d931e3148000a1c10
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
