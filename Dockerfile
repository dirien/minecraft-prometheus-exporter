# Dockerfile
FROM cgr.dev/chainguard/static@sha256:676e989769aa9a5254fbfe14abb698804674b91c4d574bb33368d87930c5c472
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
