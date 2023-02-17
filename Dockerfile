# Dockerfile
FROM cgr.dev/chainguard/static@sha256:d9dd790fb308621ac4a5d648a852fbc455cda12f487eb30fb775a479c4f90703
COPY minecraft-exporter \
	/usr/bin/minecraft-exporter
ENTRYPOINT ["/usr/bin/minecraft-exporter"]
