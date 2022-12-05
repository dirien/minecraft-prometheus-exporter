# Dockerfile
FROM cgr.dev/chainguard/static@sha256:4200823b7d17c72aece95ead78940c1c15d2e575180a1ee0008cc62a051a584d
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
