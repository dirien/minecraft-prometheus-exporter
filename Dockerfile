# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d4c20db9cb2dbf1ac9ec77f9dbc11080a78514a5f9b96096965550dbd1c73e09
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
