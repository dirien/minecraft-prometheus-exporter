# Dockerfile
FROM cgr.dev/chainguard/static@sha256:b4437b52c1c59f8376bbf51dcf98ea74211b9ab3e62c2709dfb6518cd69467c4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
