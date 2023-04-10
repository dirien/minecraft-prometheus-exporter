# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6adc5153923b90f42f037be381bc542b29b5103f50a14b7326eb7027062cec25
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
