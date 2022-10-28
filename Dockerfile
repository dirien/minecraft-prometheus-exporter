# Dockerfile
FROM cgr.dev/chainguard/static@sha256:fa009e98c5e48c7b4cbf058a915a23ccab01e056344ebce58b52d0b62d284ec2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
