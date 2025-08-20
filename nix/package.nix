{
  version,
  lib,
  buildGoModule,
}:
buildGoModule {
  name = "api-acmcsuf";
  src = ../.;
  version = version;
  vendorHash = "sha256-3BpLbfoLmv5dwBxEHW7i//MlgIsrGMT8ZILrze7WM18=";
  postBuild = ''
    mv $GOPATH/bin/api $GOPATH/bin/api-acmcsuf
  '';

  subPackages = ["cmd/api" "cmd/csuf"];

  meta = {
    description = "API created and used by CSUF's ACM chapter";
    homepage = "https://github.com/acmcsufoss/api.acmcsuf.com";
    license = lib.licenses.mit;
  };
}
