# Dockerfile
FROM cgr.dev/chainguard/static@sha256:90949231aa480f37c8a689ed08965aaec0ba0fbc684453890f722feb8ecceb4a
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
