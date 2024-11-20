# Dockerfile
FROM cgr.dev/chainguard/static@sha256:5ff428f8a48241b93a4174dbbc135a4ffb2381a9e10bdbbc5b9db145645886d5
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
