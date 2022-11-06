# Dockerfile
FROM cgr.dev/chainguard/static@sha256:0c627e7a97a91ed2049c3da50a915b8577a79821561a5704ec89f6c8972deae2
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
