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
        postgresDir = toString ./.;
      in
      {
        devShells.default = with pkgs; mkShell {
          buildInputs = [
            glibcLocales
            couchdb3
            redis
            postgresql
            nginx
            bun
            zsh
          ];

         
          shellHook = ''
            
            # set .zshrc
            cat << EOF > .zshrc

            ### locale 
            export LOCALE_ARCHIVE=${pkgs.glibcLocales}/lib/locale/locale-archive

            ### couchdb
            export ERL_FLAGS="-couch_ini $PWD/.couchdb/config/local.ini"

            ### redis 
            # set overcommit_memory
            if [[ $(sysctl -n vm.overcommit_memory) -eq 0 ]]; then
              sudo sysctl -w vm.overcommit_memory=1
            fi

            ### postgresql 
            export PGDATA=$PWD/.postgres/data       
            export PGHOST=$PWD/.postgres/socket      

            mkdir -p \$PGDATA
            mkdir -p \$PGHOST 

            # Initialize PostgreSQL if needed
            if [ ! -f "\$PGDATA/PG_VERSION" ]; then
              initdb -D \$PGDATA
            fi

            ### shell prompt
            export PS1="%F{green}(JW):%F{blue}%c%F{white}$ "
            EOF

            ### zsh
            export SHELL=$(which zsh)
            export ZDOTDIR=$PWD
            exec zsh
          '';
        };
      }
    );
}
