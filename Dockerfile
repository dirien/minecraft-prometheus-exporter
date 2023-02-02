# Dockerfile
FROM cgr.dev/chainguard/static@sha256:922df171b70559350055c695693122f3162ef2eb337ec34ae59cd0cb13042680
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
