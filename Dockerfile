# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e78eb21f59f52446a23f7d45e78805c4f26406a6e9c4b21ec50dfdf07e6bec57
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
