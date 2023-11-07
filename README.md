# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services!

## Development

### Code generation

Generate all the code in this project by running the following command:

```sh
go generate ./...
```

### Run

To run the server, run the following command:

```sh
go run .
```

### Format

Format the source code by running the following command:

```sh
go fmt ./...
```

### Nil panic detection

To detect nil panics, run the following command:

```sh
go install go.uber.org/nilaway/cmd/nilaway@latest
```

```sh
nilaway ./...
```

### Architecture

Diamond's suggestion outline of the project structure of a Go project that generates sqlc->go and jsonnet->openapi->go.

```
- storage/
    - sqlc.go
    - postgres/
        - schema.sql
        - queries.sql
    - sqlite/
        - schema.sql
        - queries.sql
- server/
    - server.go
    - routes/
- openapi/
    - openapi.jsonnet
    - lib/
```
