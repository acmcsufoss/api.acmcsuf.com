{
  description = "Generic Dev Environment";

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
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gotools
            gopls # Go langauge server
            nilaway # Go static analysis tool
            sqlc # compiles SQL queries to Go code
            sqlfluff # SQL linter
            gnumake
          ];

          shellHook = ''
            echo "Loaded dev shell."
          '';
        };
      }
    );
}
