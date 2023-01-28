# Dockerfile
FROM cgr.dev/chainguard/static@sha256:1d476e58f42dac0abc19145a2146ae791653c7f79729f2f4c75b3fd073065797
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
