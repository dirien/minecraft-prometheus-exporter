# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d3041c12f55d84bb1477e76cdbab3fce2ac93406671386dd36fc9120b39f4da9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
