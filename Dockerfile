# Dockerfile
FROM cgr.dev/chainguard/static@sha256:fd59d10894f38ce93eb6e587595ccdd8570bfd9c8f6fde7df4c589a5cefd82e2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
