{
  description = "Go Dev Environment for api.acmcsuf.com";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages.default = pkgs.buildGoModule {
          name = "api-acmcsuf";
          src = ./.;
          vendorHash = "sha256-xcJrcUGnMKkYxJnNsJ3nvSWHb8u2ujRv/PmJMJdEWhM=";

          nativeBuildInputs = with pkgs; [
            sqlc
          ];

          env.CGO_ENABLED = 0;

          preBuild = ''
            go generate ./...
          '';

          postBuild = ''
            mv $GOPATH/bin/api $GOPATH/bin/api-acmcsuf
          '';

          subPackages = ["cmd/api"];

          meta = with pkgs.lib; {
            description = "api for acmcsuf oss";
            homepage = "https://github.com/acmcsufoss/api.acmcsuf.com";
            license = licenses.mit;
          };
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gotools
            gopls # Go langauge server
            nilaway # Go static analysis tool
            delve # Go debugger
            sqlc # compiles SQL queries to Go code
            sqlfluff # SQL linter
            gnumake
          ];

          shellHook = ''
            export DATABASE_URL="file:dev.db?cache=shared&mode=rwc"
            export CGO_ENABLED=0  # cgo compiler flags cause issues with delve when using Nix
            echo "Loaded dev shell."
          '';
        };
      }
    );
}
