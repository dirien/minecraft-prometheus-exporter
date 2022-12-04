# Dockerfile
FROM cgr.dev/chainguard/static@sha256:02396af5c2a2f5e986adc8bc67b1d552a983d29d4e60141b262703d3a8535a33
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
