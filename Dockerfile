# Dockerfile
FROM cgr.dev/chainguard/static@sha256:bc17ae9265675e98c7bc99b91888d1606698638418eb3fc1aeeb19110d29b300
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
