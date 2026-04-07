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
  sqlite,
  sqlite-web,
  isCI ? false,
  full ? false,
}: let
in
  mkShell {
    packages =
      [
        go
        gopls # Go language server
        go-tools
        sqlc # compiles SQL queries to Go code
        sqlfluff # SQL linter
        gnumake
        go-swag
      ]
      # Dev tools not required in CI go here
      ++ lib.optionals (!isCI) [
        air # run dev server with hot reload
        xh
        sqlite
        cobra-cli
        delve # Go debugger
      ]
      # Heavyweight or rarely used tools go here
      ++ lib.optionals full [
        sqlite-web
      ];

    env = {
      GOROOT = "${go}/share/go";
    };

    shellHook = lib.optionalString (!isCI) ''
      if [ ! -f .env ]; then
        echo ".env file not found! Creating one from .env.example for you..."
        cp .env.example .env
      fi
      echo -e "\e[32mLoaded nix dev shell\e[0m"
    '';
  }
