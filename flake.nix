{
  description = "Go Dev Environment for api.acmcsuf.com";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.11";
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

        packages = import ./nix/package.nix {
          inherit pkgs;
        };
        devShells = import ./nix/devShell.nix {
          inherit pkgs;
        };
      in {
        packages.default = packages.default;
        devShells.default = devShells.default;
      }
    );
}
