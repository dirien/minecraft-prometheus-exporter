# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0394eb00eabee3530bdfc0dd5845806e5badcc840cbd0746b0d5bd912c9d51fc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
