# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c7cc7496526ad868bfdb0411544204d25a4a2a748d536ac9ae086fb04a2b08fc
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
