# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services! See the `developer-docs/` directory for more information on using and contributing to this project.

---

This project requires that you have Go, sqlc, GNU Make, and optionally Air installed. We recommend using the
provided Nix development environment, see below.

### Start API server
Air will recompile the project on the fly so you don't have to restart the server when you make changes.
```sh
air
```

To compile and run manually, you can use one of the following:
```sh
make # Compile program
./bin/api # Run program
```
OR
```sh
make run
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

1. [Install nix](https://determinate.systems/nix-installer/) and optionally [direnv](https://direnv.net/docs/installation.html) if you don't already have them

2. Run `direnv allow` at the project root. If you don't want to use direnv, you can use `nix develop` to achieve the same thing, but you will need to run it every time you enter the project.

Developed with 💚 by [**@acmcsufoss**](https://github.com/acmcsufoss)
