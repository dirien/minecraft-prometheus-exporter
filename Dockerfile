# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c423ec4109b1851ee1a0fb6d8f785c73841e43757d320cd03eae2869e9144022
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
