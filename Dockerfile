# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d7518504f59dacbc90852349c0878871199cefd4bed7952d2eeb7cc3ddbe69e5
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
