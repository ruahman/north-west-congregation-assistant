
couchdb\:start:
	couchdb

redis\:start:
	redis-server

postgres\:start:
	@echo start postgres

nginx\:start:
	@echo start nginx

clean\:postgres:
	rm -rf .postgres

clean\:couchdb:
	rm -rf .couchdb

clean\:redis:
	rm -rf .redis

clean\:zshrc:
	rm .zshrc

clean: clean\:couchdb clean\:postgres clean\:redis clean\:zshrc

.PHONEY: couchdb\:start redis\:start postgres\:start nginx\:start
.PHONEY: clear clear\:postgres clear\:couchdb clean\:redis clean\:zshrc

