# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services!

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
