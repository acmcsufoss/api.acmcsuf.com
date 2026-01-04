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
  bash,
  bash-completion,
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
    bash
    bash-completion
    ];

  shellHook = ''
  # Shell auto complete
  source ${bash-completion}/etc/profile.d/bash_completion.sh

  if [ ! -d bin ]; then 
	  mkdir bin # Nix blows up if bin isnt here 
  fi

  COMPLETION_FILE=./bin/completion
  if [ -f "$COMPLETION_FILE" ]; then 
  	source "$COMPLETION_FILE"
  else
  	echo "Note: Your cli has no autocomplete, remember to source completion in ./bin or restart the shell when it is generated!"
  fi
  

    export CGO_ENABLED=0  # cgo compiler flags cause issues with delve when using Nix
    if [ ! -f .env ]; then
      echo ".env file not found! Creating one from .env.example for you..."
      cp .env.example .env
    fi
    echo -e "\e[32mLoaded nix dev shell\e[0m"
    export GOROOT="${go}/share/go"
  '';
}
