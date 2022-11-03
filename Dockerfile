# Dockerfile
FROM cgr.dev/chainguard/static@sha256:cd526a2873df9d2f5156538996d8b5b082781f62b835f4a2bbf34e5321f1cc62
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
