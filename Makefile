
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

clean: clean\:couchdb clean\:postgres

.PHONEY: couchdb\:run redis\:start postgres\:start nginx\:start
.PHONEY: clear clear\:postgres clear\:couchdb

