# Useful docker commands

```sh
# build
docker build -t acmcsuf-api:latest .

# run
docker run --rm \
--name acmcsuf-api \
--mount type=volume,source=sqlite_data,target=/app/data \
-p 8080:8080 \
acmcsuf-api:latest

```
