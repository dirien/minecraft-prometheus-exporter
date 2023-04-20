# Dockerfile
FROM cgr.dev/chainguard/static@sha256:c76fb8ae37675545fbc4335f60121ac0691b56e7d680a550d15e6c4c247b78d7
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
