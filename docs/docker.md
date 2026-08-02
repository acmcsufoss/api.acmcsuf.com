```
docker run --rm \
--name acmcsuf-api \
--mount type=volume,source=sqlite-data,target=/app/data \
-p 8080:80 \
acmcsuf-api:latest

```
