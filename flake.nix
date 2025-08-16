{
  description = "Go Dev Environment for api.acmcsuf.com";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
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
        package = pkgs.callPackage ./nix/package.nix {};
      in {
        packages.default = package;
        devShells.default = pkgs.callPackage ./nix/devShell.nix {};

        apps = {
          default = {
            type = "app";
            program = "${package}/bin/api-acmcsuf";
          };
          csuf = {
            type = "app";
            program = "${package}/bin/csuf";
          };
        };
      }
    );
}
