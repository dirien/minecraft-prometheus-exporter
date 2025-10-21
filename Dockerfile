# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bf076ce7861fe5cd50414b8ef26af247df58af0e256e17a7e4fc5ef2450393f9
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
