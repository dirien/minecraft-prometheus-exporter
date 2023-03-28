# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f274f860b7f71407279c777eb5733c57fafcc9dfc266cf2a48e08bd7fe8aee56
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
