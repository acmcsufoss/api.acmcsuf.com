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

        version =
          if (self ? rev)
          then builtins.substring 0 8 self.rev
          else "dev";

        package = pkgs.callPackage ./nix/package.nix {
          version = version;
        };
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
