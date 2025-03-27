# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services!

---

This project requires that you have Go, sqlc, and GNU Make installed. We recommend using the
provided Nix development environment, see below.

### Start API server

```sh
make # Compile program
./bin/api # Run program
```
OR
```sh
make run # This compiles & runs the program without creating a binary
```

### Other useful commands from the Makefile

```sh
make check # Run checks
make test # Run tests (None yet)
make sql-fix # Format and fix SQL files
make clean # Removes all build artifacts
```

---

## To use the Nix dev shell

1. [Install nix](https://determinate.systems/nix-installer/) and [direnv](https://direnv.net/docs/installation.html) if you don't already have them

2. Run `direnv allow` at the project root. If you don't want to use direnv, you can use `nix develop` to achieve the same thing, but you will need to run it every time you enter the project.

Developed with 💚 by [**@acmcsufoss**](https://github.com/acmcsufoss)
