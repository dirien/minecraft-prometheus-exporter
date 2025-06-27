# Dockerfile
FROM cgr.dev/chainguard/static@sha256:092aad9f6448695b6e20333a8faa93fe3637bcf4e88aa804b8f01545eaf288bd
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
