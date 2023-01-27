# Dockerfile
FROM cgr.dev/chainguard/static@sha256:285da107d4d7d8651842e512786abe7ebcf12fe867de6c8dea83958fd0ced552
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
