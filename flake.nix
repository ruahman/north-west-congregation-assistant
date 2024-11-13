{
  description = "Bun dev environment";

  inputs = {
    nixpkgs.url      = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url  = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = with pkgs; mkShell {
          buildInputs = [
            couchdb3
            nginx
            redis
            bun
          ];
         
          shellHook = ''
            ### couchdb
            export ERL_FLAGS="-couch_ini $PWD/.couchdb/config/local.ini"

            ### zsh
            export SHELL=$(which zsh)
            exec $SHELL
          '';
        };
      }
    );
}
