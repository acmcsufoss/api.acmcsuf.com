# Reference:
# https://github.com/diamondburned/acmregister/blob/2ad96e84ed069c554698261530cbd875cafe4a26/shell.nix

{ sysPkgs ? import <nixpkgs> {} }:

let
	lib = sysPkgs.lib;
	overlay = self: super: {
		go = super.go_1_21;
	};

	pkgs = import (sysPkgs.fetchFromGitHub {
		owner = "NixOS";
		repo = "nixpkgs";
		# most recent commit as of 11/13/23
		rev = "bb142a6838c823a2cd8235e1d0dd3641e7dcfd63";
		hash = "sha256:0nbicig1zps3sbk7krhznay73nxr049hgpwyl57qnrbb0byzq9iy";
	}) {
		overlays = [ overlay ];
	};
	
	nilaway = pkgs.buildGoModule rec {
		name = "nilaway";
		# version = "0.0.0-20231031155528-970a3b8379e3";

		src = pkgs.fetchFromGitHub {
			owner = "uber-go";
			repo = "nilaway";
			# rev = "v${version}";
			rev = "970a3b8379e324d48c70d34f03f3f309432b4d41";
			sha256 = "1chahgba7fqr3jz7dcqb3paa8v0xgy2qsvjwya9gqqwdzli6rdsg";
		};
		# vendorSha256 = lib.fakeSha256;
		vendorSha256 = "sha256-E44rokDNcyc8XSdvAQ/DEyltwO6zaErckAH5+JEXxrM=";
		proxyVendor = true;
		doCheck = false;
	};
in pkgs.mkShell {
	buildInputs = with pkgs; [
		go
		gotools
		gopls
		sqlc
		nilaway
		# TODO: Add SQL formatting tool.
	];
}

