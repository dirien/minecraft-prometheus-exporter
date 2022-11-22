# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5ba44445d663b5d1b9e844726ab2c85fcf3fc68662612a3674aab17d484fecc4
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
