Build image:
```sh
docker build -t acmcsuf-api:latest
```

Remove image:
```sh
docker rmi acmcsuf-api
```

Create volume for database:
```sh
docker volume create acmcsuf-data
```

Run:
```sh
docker run -p 8080:80 -v acmcsuf-data:/app/data acmcsuf-api:latest
```
