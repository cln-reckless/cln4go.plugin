{
  description = "A simple core lightning plugin that semplify the Splice operation";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    goflake.url = "github:sagikazarmark/go-flake";
    goflake.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils, goflake, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages = {
          default = pkgs.gnumake;
        };
        formatter = pkgs.nixpkgs-fmt;

        devShell = pkgs.mkShell {
          hardeningDisable = [ "fortify" ];
          buildInputs = with pkgs; [ gnumake golangci-lint libcap go gcc ];
          buildPhase = "make";
          shellHook = ''
            export CGO_ENABLED=0

            make dep
          '';
        };
      }
  );
}
