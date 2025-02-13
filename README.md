# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services!

## Develop

### Start API server

```sh
go run cmd/api/main.go
```

### Generate code

```sh
go generate ./...
```

### Run tests

```sh
go test ./...
```

### Format code

```sh
go fmt ./...
```

---

## Or use the Makefile instead

### Start API server

```sh
make run
```

### Run checks and tests

```sh
make test
```

### Format and fix SQL files

```sh
make sql-fix
```

---

## To use the Nix dev shell

1. [Install nix](https://determinate.systems/nix-installer/) and [direnv](https://direnv.net/docs/installation.html) if you don't already have them

2. Run `direnv allow` at the project root. If you don't want to use direnv, you can use `nix develop` to achieve the same thing, but you will need to run it every time you enter the project.

Developed with 💚 by [**@acmcsufoss**](https://github.com/acmcsufoss)
