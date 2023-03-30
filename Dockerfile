# Dockerfile
FROM cgr.dev/chainguard/static@sha256:39b1975494f45bf36eb87f08ce6e90dd633c91b10fd783d824a5a3836f0d6fdc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
