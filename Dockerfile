# Dockerfile
FROM cgr.dev/chainguard/static@sha256:29fb66034dd7174ddd2538a69b84127db1b072f3f7714d9f3a8f8d598ef19ce9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
