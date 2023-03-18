# Dockerfile
FROM cgr.dev/chainguard/static@sha256:03f490274beb5b3c78044b4451bac6d39c68b25e12c568fe6274af334b122b2a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
