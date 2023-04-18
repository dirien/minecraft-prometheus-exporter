# Dockerfile
FROM cgr.dev/chainguard/static@sha256:f0b465f9fff39eb309f1e7ace4e11900ccba5a93bdb51c3481d1d79d5081a690
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
