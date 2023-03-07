# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c753c13161f19921c8eca4438dd087746501de26363424f0d9d10470a35ad927
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
