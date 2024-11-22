default: nix

nix:
	nix develop

couchdb\:start:
	@echo start couchdb...
	couchdb

couchdb\:clean:
	rm -rf .config/couchdb

redis\:start:
	@echo start redis...
	redis-server $$REDISDIR/redis.conf

redis\:clean:
	rm -rf .config/redis

postgres\:start:
	@echo start postgres...
	pg_ctl -D $$PGDATA -o "-k $$PGDIR" -l $$PGDIR/logfile start

postgres\:clean:
	rm -rf .config/postgres

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

nginx\:reload:
	@echo reload nginx...
	nginx -e stderr -c $$NGINXDIR/nginx.conf -s reload

nginx\:clean:
	rm -rf .config/nginx

.PHONY: zshrc\:clean
zshrc\:clean:
	rm .config/.zshrc

.PHONY: node\:clean
node\:clean:
	rm -rf node_modules

.PHONY: clean
clean: couchdb\:clean postgres\:clean redis\:clean nginx\:clean zshrc\:clean
	rm -rf .config

.PHONY: nix
.PHONY: couchdb\:start couchdb\:clean 
.PHONY: redis\:start redis\:clean
.PHONY: postgres\:start postgres\:stop postgres\:clean psql\:connect psql
.PHONY: nginx\:start nginx\:stop nginx\:reload nginx\:clean 

