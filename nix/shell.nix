{
  mkShell,
  go,
  gopls,
  delve,
  sqlc,
  air,
  sqlfluff,
  gnumake,
  curl,
  xh,
  jq,
  go-swag,
  cobra-cli,
  go-tools,
}:
mkShell {
  packages = [
    go
    gopls # Go language server
    go-tools
    delve # Go debugger
    sqlc # compiles SQL queries to Go code
    air # run dev server with hot reload
    sqlfluff # SQL linter
    gnumake
    curl
    xh
    jq
    go-swag
    cobra-cli
  ];

  shellHook = ''
    export CGO_ENABLED=0  # cgo compiler flags cause issues with delve when using Nix
    if [ ! -f .env ]; then
      echo ".env file not found! Creating one from .env.example for you..."
      cp .env.example .env
    fi
    echo -e "\e[32mLoaded nix dev shell\e[0m"
  '';
}
