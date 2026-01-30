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
  go-migrate,
}:

let
	go-migrate-sqlite = go-migrate.overrideAttrs (oldAttrs: {
		tags = [ "sqlite3" ];
	});
in
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
    go-migrate-sqlite
  ];

  shellHook = ''
    export CGO_ENABLED=0  # cgo compiler flags cause issues with delve when using Nix
    if [ ! -f .env ]; then
      echo ".env file not found! Creating one from .env.example for you..."
      cp .env.example .env
    fi
    if [ -d "./migrations" ]; then 
    	if ! find "./migrations" -mindepth 1 -maxdepth 1 | read; then
		echo "CREATING MIGRATIONS"
		migrate create -ext sql -dir migrations init
	fi
    else
    	echo "CREATING MIGRATIONS"
    	mkdir migrations
	migrate create -ext sql -dir migrations init
    fi
    echo -e "\e[32mLoaded nix dev shell\e[0m"
    export GOROOT="${go}/share/go"
  '';
}
