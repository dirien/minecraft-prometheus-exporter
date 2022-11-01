# Dockerfile
FROM cgr.dev/chainguard/static@sha256:44919ee2b6fa7c2d435e5ffa6de1f8cf4ebb8e89c4f02fcda2cc088e4550c728
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
