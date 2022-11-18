# Dockerfile
FROM cgr.dev/chainguard/static@sha256:e1d2d18b91ad46e8098876135cb38844f20636eb1fd4d2ae0942c44845568bfa
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
