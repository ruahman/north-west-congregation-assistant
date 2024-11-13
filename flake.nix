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
            glibcLocales
            postgresql
            couchdb3
            redis
            nginx
            bun
          ];

         
          shellHook = ''
            ### locale 
            export LOCALE_ARCHIVE=${pkgs.glibcLocales}/lib/locale/locale-archive

            ### postgresql 

            ### couchdb
            export ERL_FLAGS="-couch_ini $PWD/.couchdb/config/local.ini"

            ### redis 
            sudo sysctl -w vm.overcommit_memory=1

            ### nginx

            ### zsh
            # export SHELL=$(which zsh)
            # exec $SHELL
          '';
        };
      }
    );
}
