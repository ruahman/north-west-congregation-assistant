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

        COUCHDB_DIR = "$PWD/.config/couchdb";

        local_ini = ''
        [couchdb]
        database_dir = ${COUCHDB_DIR}/data

        [admins]
        admin = ruahman

        [httpd]
        bind_address = 127.0.0.1
        '';

        REDIS_DIR = "$PWD/.config/redis";

        redis_conf = ''
        dir ${REDIS_DIR} 

        dbfilename dump.rdb
        '';

        POSTGRES_DIR = "$PWD/.config/postgres";

        NGINX_DIR = "$PWD/.config/nginx";

        index_html = ''
        <h1>Hello Territory Assistant</h1>
        '';

      in
      {
        devShells.default = with pkgs; mkShell {
          buildInputs = [
            openssl
            glibcLocales
            couchdb3
            redis
            postgresql
            nginx
            typescript-language-server
            bun
            zsh
          ];
         
          shellHook = ''
            
            export LOCALE_ARCHIVE=${pkgs.glibcLocales}/lib/locale/locale-archive

            ### couchdb
            mkdir -p ${COUCHDB_DIR}/data

            if [ ! -f "${COUCHDB_DIR}/local.ini" ]; then
              echo "${local_ini}" > "${COUCHDB_DIR}/local.ini"
            fi

            ### redis 
            mkdir -p ${REDIS_DIR}

            if [ ! -f "${REDIS_DIR}/redis.conf" ]; then
              echo "${redis_conf}" > "${REDIS_DIR}/redis.conf"
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
            mkdir -p ${NGINX_DIR}/certs

            if [ ! -f "${NGINX_DIR}/nginx.conf" ]; then
              cat << EOF > "${NGINX_DIR}/nginx.conf"
            # setup worker processes according to cores you have
            worker_processes auto;

            # how many connections each worker processes can handle
            events {
                worker_connections 1024; 
            }

            pid ${NGINX_DIR}/logs/nginx.pid;  # set custom pid file location

            # handle all http requests
            http {
                # set custom log paths
                error_log  ${NGINX_DIR}/logs/error.log;
                access_log ${NGINX_DIR}/logs/access.log;
    
                types {
                    text/html                             html htm shtml;
                    text/css                              css;
                    text/xml                              xml;
                    image/gif                             gif;
                    image/jpeg                            jpeg jpg;
                    application/javascript                js;
                    application/json                      json;
                    image/png                             png;
                    image/svg+xml                         svg svgz;
                    video/mp4                             mp4;
                    video/webm                            webm;
                    audio/mpeg                            mp3;
                    audio/ogg                             ogg;
                    font/woff                             woff;
                    font/woff2                            woff2;
                    application/octet-stream              bin exe;
                    application/pdf                       pdf;
                    application/zip                       zip;
                    application/x-gzip                    gz;
                }

                # list of backend server that we can forward to
                # by default it does a round robin
                # upstream backendservers {
                #     server 127.0.0.1:1111;
                #     server 127.0.0.1:1112;
                #     server 127.0.0.1:1113;
                #     server 127.0.0.1:1114;
                # }

                # basic http server
                server {
                    listen       8080;            # port number to listen on
                    server_name  localhost;       # server name (can be a domain or localhost)
                    
                    # redirect all http to https
                    # return 301 https://$host$request_uri;

                    # serve files from the 'html' directory
                    root   ${NGINX_DIR}/html;                  # path to your static files
                    index  index.html;            # default file to serve

                    # rewrite path not redirect
                    # rewrite ^/number/(\w+) /count/$1;

                    # define location block
                    location / {
                        # try finding file path in \$uri,
                        # if not fourn try \$uri/
                        # if still not found return =404
                        try_files \$uri \$uri/ =404;
                    }

                    # proxy to just one proxy server 
                    # location / {
                    #     proxy_pass http://backend_server_address;
                    #     proxy_set_header Host $host;
                    #     proxy_set_header X-Real-IP $remote_addr; 
                    #     ....
                    # }

                    # proxy to backendservers 
                    # location / {
                    #     proxy_pass http://backendservers/;
                    # }

                    # append assets to root path
                    # location /assets/ {
                    #     root ${NGINX_DIR}/html;  
                    # }

                    # alias for html/fruits
                    # location /carbs {
                    #     alias ${NGINX_DIR}/html/fruits;  # alias for html/fruits
                    # }

                    # regular expression location block
                    # location ~* /count/[0-9] {
                    #     root ${NGINX_DIR}/html;
                    #     try_files /index_html =404;
                    # }

                    # redirect 
                    # location /crops {
                    #   return 307 /fruits;
                    # }
                }
                
                # https server
                # server {
                #     listen 443 ssl;
                #     server_name example.com www.example.com;
                #
                #     # SSL Configuration 
                #     ssl_certificate ${NGINX_DIR}/cert/pub.cert;
                #     ssl_certificate_key ${NGINX_DIR}/cert/priv.key;;
                #
                #     # security headers 
                #     add_header Strict-Transport-Securty "max-age=2345235; includeSubDomains" always;
                #
                #     location / {
                #         root /foo/bar;
                #         index index.html;
                #     }
                #
                # }
            }
            EOF
            fi

            if [ ! -f "${NGINX_DIR}/html/index.html" ]; then
              echo "${index_html}" > "${NGINX_DIR}/html/index.html"
            fi

            if [ ! -f "${NGINX_DIR}/certs/priv.key" ]; then
              cd "${NGINX_DIR}/certs"
              # -x509: cert type
              # -nodes: no passcode 
              # -days: expiration 
              # -newkey: algorythm
              openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout priv.key -out pub.crt
              cd ../../..
            fi

            ### zsh
            cat << EOF > .config/.zshrc

            export LOCALE_ARCHIVE=${pkgs.glibcLocales}/lib/locale/locale-archive
            export COUCHDBDIR=${COUCHDB_DIR}
            export ERL_FLAGS="-couch_ini \$COUCHDBDIR/local.ini"
            export REDISDIR=${REDIS_DIR}
            export PGDIR=${POSTGRES_DIR}      
            export PGDATA=\$PGDIR/data       
            export NGINXDIR=${NGINX_DIR}

            git_status() {
  
              # Check if inside a Git repository
              if ! git rev-parse --is-inside-work-tree &> /dev/null; then
                  return 1  # Exit the function early with a non-zero exit code
              fi

              # get current branch
              branch=\$(git branch --show-current)

              # check if any files were modified
              if [[ -n \$(git status --porcelain | grep '^ \?M') ]]; then
                changes="*"
              else
                changes=""
              fi

              # Check if any files were added
              if [[ -n \$(git status --porcelain | grep '^ \?A') ]]; then
                added="+"
              else
                added=""
              fi

              # check if any files were deleted
              if [[ -n \$(git status --porcelain | grep '^ \?D') ]]; then
                deleted="-"
              else
                deleted=""
              fi

              # check if there are any untracked files
              if [[ -n \$(git status --porcelain | grep '^ \??') ]]; then
                untracked="?"
              else
                untracked=""
              fi

              # check if your branch is ahead
              if git status | grep -q "Your branch is ahead"; then
                ahead="" 
              else
                ahead=""
              fi

              echo " %F{yellow}\$changes\$added\$deleted\$(echo \$untracked)%F{red}git(\$branch%F{yellow}\$ahead%F{red})"
            } 
  
            export start_path=\$PWD
            relative_path() {
              if [[ -z \$1 ]]; then
                  return 1
              fi

              # Get the relative path using realpath
              if relative=\$(realpath --relative-to="\$(echo \$start_path)" "\$1" 2>/dev/null); then
                  if [[ "\$relative" == "." ]]; then
                      echo "~"
                  else
                      echo "~/\$relative"
                  fi
              else
                  return 1
              fi
            }

            # Define a function to generate the prompt
            function update_prompt {
                PROMPT="%F{green}󱄅 (%n@%M):%F{blue}\$(relative_path \$PWD)\$(git_status)%F{white}$ "
                export PS1=\$PROMPT
            }

            precmd_functions+=(update_prompt)
            EOF

            if [ -t 1 ]; then
              export SHELL=$(which zsh)
              export ZDOTDIR=$PWD/.config
              exec zsh
            fi
          '';
        };
      }
    );
}
