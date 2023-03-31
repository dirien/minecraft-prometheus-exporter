# Dockerfile
FROM cgr.dev/chainguard/static@sha256:71e62a091a9b7314b6ba2f5313fc59778ed79951c534af7ecc3e5f974a994cf2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
