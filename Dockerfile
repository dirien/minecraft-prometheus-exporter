# Dockerfile
FROM cgr.dev/chainguard/static@sha256:a53dc4ceb82c6758f2d746a96474f15d10b72d70584e45a7a9f4975aca69d1e9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
