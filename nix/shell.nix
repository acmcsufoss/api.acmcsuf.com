{
  mkShell,
  lib,
  go,
  gopls,
  delve,
  sqlc,
  air,
  sqlfluff,
  gnumake,
  xh,
  go-swag,
  cobra-cli,
  go-tools,
  go-migrate,
  sqlite,
  sqlite-web,
  full ? false,
}: let
  go-migrate-sqlite = go-migrate.overrideAttrs (oldAttrs: {
    tags = ["sqlite3"];
  });
in
  mkShell {
    packages =
      [
        go
        gopls # Go language server
        go-tools
        delve # Go debugger
        sqlc # compiles SQL queries to Go code
        air # run dev server with hot reload
        sqlfluff # SQL linter
        gnumake
        xh
        go-swag
        cobra-cli
        go-migrate-sqlite
        sqlite
      ]
      ++ lib.optionals full [
        # Heavyweight/optional dependencies
        sqlite-web
      ];

    shellHook = ''
      if [ ! -f .env ]; then
        echo ".env file not found! Creating one from .env.example for you..."
        cp .env.example .env
      fi
      echo -e "\e[32mLoaded nix dev shell\e[0m"
    '';
  }
