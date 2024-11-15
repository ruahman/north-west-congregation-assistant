{
  description = "Territory Assitant dev environment";

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

        REDIS_DIR = "$PWD/.redis";

        POSTGRES_DIR = "$PWD/.postgres";

        NGINX_DIR = "$PWD/.nginx";

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
            
            export LOCALE_ARCHIVE=${pkgs.glibcLocales}/lib/locale/locale-archive

            ### couchdb
            mkdir -p ${COUCHDB_DIR}/data

            if [ ! -f "${COUCHDB_DIR}/local.ini" ]; then
              cat << EOF > "${COUCHDB_DIR}/local.ini" 
            [couchdb]
            database_dir = ${COUCHDB_DIR}/data

            [admins]
            admin = ruahman

            [httpd]
            bind_address = 127.0.0.1
            EOF
            fi

            ### redis 
            mkdir -p ${REDIS_DIR}
            if [ ! -f "${REDIS_DIR}/redis.conf" ]; then
              cat << EOF > "${REDIS_DIR}/redis.conf" 
            dir ${REDIS_DIR} 

            dbfilename dump.rdb
            EOF
            fi
            
            # set overcommit_memory
            if [[ $(sysctl -n vm.overcommit_memory) -eq 0 ]]; then
              sudo sysctl -w vm.overcommit_memory=1
            fi

            ### postgres 
            mkdir -p ${POSTGRES_DIR}/data

            # Initialize PostgreSQL if needed
            if [ ! -f "${POSTGRES_DIR}/data/PG_VERSION" ]; then
              initdb -D ${POSTGRES_DIR}/data
            fi

            ### nginx
            mkdir -p ${NGINX_DIR}/logs
            mkdir -p ${NGINX_DIR}/html

            if [ ! -f "${NGINX_DIR}/nginx.conf" ]; then
              cat << EOF > "${NGINX_DIR}/nginx.conf"
            # nginx.conf

            # define the global settings
            events {}

            pid ${NGINX_DIR}/logs/nginx.pid;  # set custom pid file location

            http {
                # set custom log paths
                error_log  ${NGINX_DIR}/logs/error.log;
                access_log ${NGINX_DIR}/logs/access.log;

                # server block
                server {
                    listen       8080;            # port number to listen on
                    server_name  localhost;       # server name (can be a domain or localhost)

                    # serve files from the 'html' directory
                    root   ${NGINX_DIR}/html;                  # path to your static files
                    index  index.html;            # default file to serve

                    # define location block
                    location / {
                        try_files \$uri \$uri/ =404;
                    }

                    # Optionally serve other static directories
                    location /assets/ {
                        root html;                # Serve assets from 'html/assets'
                    }
                }
            }
            EOF
            fi

            if [ ! -f "${NGINX_DIR}/html/index.html" ]; then
              cat << EOF > "${NGINX_DIR}/html/index.html"
              <h1>Hello Territory Assitant</h1>
            EOF
            fi

            ### zsh
            cat << EOF > .zshrc

            export LOCALE_ARCHIVE=${pkgs.glibcLocales}/lib/locale/locale-archive
            export COUCHDBDIR=${COUCHDB_DIR}
            export ERL_FLAGS="-couch_ini \$COUCHDBDIR/local.ini"
            export REDISDIR=${REDIS_DIR}
            export PGDIR=${POSTGRES_DIR}      
            export PGDATA=\$PGDIR/data       
            export NGINXDIR=${NGINX_DIR}

            export PS1="%F{green}(JW):%F{blue}%c%F{white}$ "
            EOF

            export SHELL=$(which zsh)
            export ZDOTDIR=$PWD
            exec zsh
          '';
        };
      }
    );
}
