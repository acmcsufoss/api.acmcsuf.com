{
  version,
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "acmcsuf-api";
  src = ../.;
  version = version;
  vendorHash = "sha256-XCzKPPjCbO2Kp9XGZ94PpPedyJhFAs+ys+gna66jDSE=";
  subPackages = ["cmd/acmcsuf-api" "cmd/acmcsuf-cli"];
  ldflags = ["-X main.Version=${version}"];

  meta = {
    description = "API with CLI wrapper created and used by CSUF's ACM chapter";
    homepage = "https://github.com/acmcsufoss/api.acmcsuf.com";
    license = lib.licenses.mit;
  };
}
