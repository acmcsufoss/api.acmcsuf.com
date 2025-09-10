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
    export DATABASE_URL="file:dev.db?cache=shared&mode=rwc"
    export CGO_ENABLED=0  # cgo compiler flags cause issues with delve when using Nix
    export PATH="$PWD/bin:$PATH"
    echo -e "\e[32mLoaded nix dev shell\e[0m"
  '';
}
