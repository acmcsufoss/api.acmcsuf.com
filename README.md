# api.acmcsuf.com

acm@CSUF's API for managing events, announcements, forms, and other resources. See below for info on running/developing this project locally, and see [`CONTRIBUTING.md`](./.github/CONTRIBUTING.md) and [`developer-docs/`](./developer-docs) for more information on using and contributing to this project.
  
This project contains two *separate* but related binaries: the API server itself, and a CLI wrapper that makes interacting with the API easier.

---

## Getting Started
This project requires that you have Go, sqlc, GNU Make, and optionally Air installed. We recommend using the
provided Nix development environment (which will provide all of these and more), but it's not required.

### Option 1: With Nix
1. [Install nix](https://determinate.systems/nix-installer/) and optionally [direnv](https://direnv.net/docs/installation.html) if you don't already have them
2. Run `direnv allow` once at the project root if you have direnv. Otherwise use `nix develop` to enter the dev shell each time.
> If you're using VSCode, make sure to run `code .` AFTER entering the dev shell in your terminal.

### Option 2: Without Nix
1. Aquire dependencies:
    - Go v1.24
    - sqlc
    - GNU Make
    - Staticcheck (optional)
    - Air (optional)

2. Create a `.env` file with the following command:
```sh
cp .env.example .env
```
3. Load the environment variables with the following command:
```sh
source .env
```


## Compiling and Running the Server and CLI
Use one of the following commands to compile and run the API server:
```sh
air      # runs server with hot reloading (recommended)
make run # runs binary directly
```

Now visit <http://localhost:8080/swagger/index.html> in your browser!

### To use the CLI:
```sh
make
./bin/acmcsuf-cli # if you have direnv you can just run `acmcsuf-cli`
```

### Other Makefile Commands

```sh
make check   # Run linter
make test    # Run unit tests (none yet)
make sql-fix # Format and fix SQL files
make clean   # Removes all build artifacts
```

---

Developed with 💚 by [**@acmcsufoss**](https://github.com/acmcsufoss)  
Please create an issue or contact one of the team leads if you have any isses.  
Current team leads: [**Josh Holman**][tl_1] and [**Siddharth Vasu**][tl_2]

[tl_1]: https://github.com/thejolman
[tl_2]: https://github.com/sidvasu

