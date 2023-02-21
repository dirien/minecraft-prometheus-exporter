# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bfd9115197feb9c47f116aa7d90a295b7d535d9ef721ac9330435cac207752a8
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
