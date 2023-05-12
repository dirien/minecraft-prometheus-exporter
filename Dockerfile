# Dockerfile
FROM cgr.dev/chainguard/static@sha256:6565f8cd1399f120cebc566096d3ce51a93fef90c787ab3524b5bbdfe829f82f
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
