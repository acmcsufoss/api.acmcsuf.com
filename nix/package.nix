{
  version,
  lib,
  buildGoModule,
}:
buildGoModule {
  name = "acmcsuf-api";
  src = ../.;
  version = version;
  vendorHash = "sha256-3BpLbfoLmv5dwBxEHW7i//MlgIsrGMT8ZILrze7WM18=";
  subPackages = ["cmd/acmcsuf-api" "cmd/acmcsuf-cli"];
  ldflags = ["-X main.Version=${version}"];

  meta = {
    description = "API with CLI wrapper created and used by CSUF's ACM chapter";
    homepage = "https://github.com/acmcsufoss/api.acmcsuf.com";
    license = lib.licenses.mit;
  };
}
