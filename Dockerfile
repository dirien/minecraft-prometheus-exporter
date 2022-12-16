# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f9998530e92993931755e4a84a2078ccc2a819d516ddefb66548848f273fb5e0
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
