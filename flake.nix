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
            # first check then set
            # sudo sysctl -w vm.overcommit_memory=1

            ### nginx

            ### shell prompt
            export PS1="\e[0;32m(JW)\e[0m:\e[0;34m\W\e[0m\$ "
          '';
        };
      }
    );
}
