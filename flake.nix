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
            mkdir -p ${COUCHDB_DIR}/data
            mkdir -p ${COUCHDB_DIR}/config

            if [ ! -f "${COUCHDB_DIR}/config/local.ini" ]; then
              echo "${local_ini}" > "${COUCHDB_DIR}/config/local.ini"
            fi

            export ERL_FLAGS="-couch_ini $PWD/.couchdb/config/local.ini"

            ### redis 
            mkdir -p ${REDIS_DIR}/data
            mkdir -p ${REDIS_DIR}/config
            if [ ! -f "${REDIS_DIR}/config/redis.conf" ]; then
              echo "${redis_conf}" > "${REDIS_DIR}/config/redis.conf" 
            fi
            
            # set overcommit_memory
            if [[ $(sysctl -n vm.overcommit_memory) -eq 0 ]]; then
              sudo sysctl -w vm.overcommit_memory=1
            fi

            ### postgresql 
            export PGDATA=${POSTGRES_DIR}/data       
            export PGHOST=${POSTGRES_DIR}/socket      

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
