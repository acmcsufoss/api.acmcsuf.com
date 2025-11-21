# api.acmcsuf.com

ACM at CSUF club API for managing events, announcements, forms, and other services! Keep reading for information on setting up this project locally, and see the [contributor guide](./.github/CONTRIBUTING.md) and [`developer-docs/`](./developer-docs) for more information on using and contributing to this project.

---

## Setting Up
This project requires that you have Go, sqlc, GNU Make, and optionally Air installed. We highly
recommend using the provided Nix development environment instead of installing everything manually.

1. [Install nix](https://determinate.systems/nix-installer/) and [direnv](https://direnv.net/docs/installation.html) if you don't already have them

2. Run `direnv allow` at the project root
> If you don't have direnv, you can also use `nix develop` to enter the dev
> shell, but your environment variables won't get sourced automatically.

## Developing

### Compiling
Using `make` will compile both the API and CLI binaries, located at
`bin/acmcsuf-api` and `bin/acmcsuf-cli` respectively. If you installed `direnv`,
both of these binaries will be automatically added to your path. You can run
them directly like:
```sh
acmcsuf-api # Run API server
acmcsuf-cli # Start CLI client (start API server before using)
```


### Start API server
Air will recompile the project on the fly so you don't have to restart the server when you make changes.
```sh
air
```
OR
```sh
make run # Compiles and runs directly (no hot-reloading)
```

#### Configuring the API
The API server is configurable via environment variables. See [`.env.example`](./.env.example) for
values you can override and configure them in you `.env` file.

### Using the CLI
The CLI is a simple command-line client for the API server. Make sure the API
server is running before using.  
  
Output of `acmcsuf-cli --help`:
```
A CLI tool to help manage the API of the CSUF ACM website

Usage:
  acmcsuf-cli [command]

Available Commands:
  announcements Manage ACM CSUF's Announcements
  completion    Generate the autocompletion script for the specified shell
  events        A command to manage events.
  help          Help about any command
  officers      A command to manage officers.

Flags:
  -h, --help      help for acmcsuf-cli
  -v, --version   version for acmcsuf-cli

Use "acmcsuf-cli [command] --help" for more information about a command.
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
