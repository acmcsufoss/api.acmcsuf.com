{pkgs}: {
  default = pkgs.buildGoModule {
    name = "api-acmcsuf";
    src = ../.;
    vendorHash = "sha256-PbObU1GpSa18kwp1kQvYTR3j+Vuh6dNLphwDYXxOejc=";

    nativeBuildInputs = with pkgs; [
      sqlc
    ];

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
}
