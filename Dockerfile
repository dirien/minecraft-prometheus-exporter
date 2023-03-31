# Dockerfile
FROM cgr.dev/chainguard/static@sha256:57a57da0f43e5acabddf65056c69dac6b049ae70335cf145e20e57c4cc32c815
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
