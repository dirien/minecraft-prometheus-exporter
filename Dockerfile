# Dockerfile
FROM cgr.dev/chainguard/static@sha256:177d0e55109c4565c5ab6fdbea232fe7fc3670b011d7dd4027f9e8a1d72f0b65
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
