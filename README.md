# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services! Keep reading for information on setting up this project locally, and see the [contributor guide](./.github/CONTRIBUTING.md) and [`developer-docs/`](./developer-docs) for more information on using and contributing to this project.

---

## Setting Up
This project requires that you have Go, sqlc, GNU Make, and optionally Air installed. We highly recommend 
using the provided Nix development environment.

1. [Install nix](https://determinate.systems/nix-installer/) and [direnv](https://direnv.net/docs/installation.html) if you don't already have them

2. Run `direnv allow` at the project root
> If you don't have direnv, you can also use `nix develop` to enter the dev
> shell, but your environment variables won't get sourced automatically.

## Developing and running

### Start API server
Air will recompile the project on the fly so you don't have to restart the server when you make changes.
```sh
air
```

To compile and run manually, you can use one of the following:
```sh
make      # Compile program
./bin/api # Run program
```
OR
```sh
make run # Compiles and runs (no hot-reloading)
```

### Other useful commands from the Makefile

```sh
make check   # Run checks
make test    # Run tests (None yet)
make fix-sql # Format and fix SQL files
make clean   # Removes all build artifacts
make help    # Display all targets
```

---

## To use the Nix dev shell


Developed with 💚 by [**@acmcsufoss**](https://github.com/acmcsufoss)
