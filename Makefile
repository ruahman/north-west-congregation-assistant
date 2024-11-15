default: nix

nix:
	nix develop

couchdb\:start:
	@echo start couchdb...
	couchdb

redis\:start:
	@echo start redis...
	redis-server $$REDISDIR/redis.conf

postgres\:start:
	@echo start postgres...
	pg_ctl -D $$PGDATA -o "-k $$PGDIR" -l $$PGDIR/logfile start

postgres\:stop:
	@echo stop postgres...
	pg_ctl -D $$PGDATA -o "-k $$PGDIR" -l $$PGDIR/logfile stop

psql:
	@echo psql... 
	psql -h $$PGDIR -p 5432 -d postgres

psql\:connect:
	@echo psql connect...
	psql -h $$PGDIR -d postgres -c '\conninfo'

nginx\:start:
	@echo start nginx...
	nginx -e stderr -q -c $$NGINXDIR/nginx.conf

nginx\:stop:
	@echo stop nginx...
	nginx -e stderr -c $$NGINXDIR/nginx.conf -s quit

clean\:postgres:
	rm -rf .postgres

clean\:couchdb:
	rm -rf .couchdb

clean\:redis:
	rm -rf .redis

clean\:nginx:
	rm -rf .nginx

clean\:zshrc:
	rm .zshrc

clean\:node_modules:
	rm -rf node_modules

clean: clean\:couchdb clean\:postgres clean\:redis clean\:nginx clean\:zshrc

.PHONY: nix
.PHONY: couchdb\:start 
.PHONY: redis\:start 
.PHONY: postgres\:start psql\:connect psql
.PHONY: nginx\:start nginx\:stop 
.PHONY: clean clean\:postgres clean\:couchdb clean\:redis clean\:zshrc clean\:node_modules clean\:nginx

