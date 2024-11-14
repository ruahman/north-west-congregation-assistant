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

        COUCHDB_DIR = "$PWD/.couchdb";

        local_ini = ''
        [couchdb]
        database_dir = $PWD/.couchdb/data

        [admins]
        admin = ruahman

        [httpd]
        bind_address = 127.0.0.1
        '';

        REDIS_DIR = "$PWD/.redis";

        redis_conf = ''
        dir $PWD/.redis 

        dbfilename dump.rdb
        '';

        POSTGRES_DIR = "$PWD/.postgres";
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
            export COUCHDBDIR=${COUCHDB_DIR}
            mkdir -p \$COUCHDBDIR/data
            mkdir -p \$COUCHDBDIR/config

            if [ ! -f "\$COUCHDBDIR/config/local.ini" ]; then
              echo "${local_ini}" > "\$COUCHDBDIR/config/local.ini"
            fi

            export ERL_FLAGS="-couch_ini \$COUCHDBDIR/config/local.ini"

            ### redis 
            export REDISDIR=${REDIS_DIR}
            mkdir -p \$REDISDIR/config
            if [ ! -f "\$REDISDIR/config/redis.conf" ]; then
              echo "${redis_conf}" > "${REDIS_DIR}/config/redis.conf" 
            fi
            
            # set overcommit_memory
            if [[ $(sysctl -n vm.overcommit_memory) -eq 0 ]]; then
              sudo sysctl -w vm.overcommit_memory=1
            fi

            ### postgresql 
            export PGDIR=${POSTGRES_DIR}      
            export PGDATA=\$PGDIR/data       

            mkdir -p \$PGDATA

            # Initialize PostgreSQL if needed
            if [ ! -f "\$PGDATA/PG_VERSION" ]; then
              initdb -D \$PGDATA
            fi

            ### shell prompt
            export PS1="%F{green}(JW):%F{blue}%c%F{white}$ "
            EOF
            # end .zshrc

            ### zsh
            export SHELL=$(which zsh)
            export ZDOTDIR=$PWD
            exec zsh
          '';
        };
      }
    );
}
