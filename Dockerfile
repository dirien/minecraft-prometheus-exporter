# Dockerfile
FROM cgr.dev/chainguard/static@sha256:8519948d4958cc565666b5b069f78390306e049dafde3b267ca3fed03e8599fa
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
