
couchdb\:start:
	couchdb

redis\:start:
	redis-server

postgres\:start:
	@echo start postgres

nginx\:start:
	@echo start nginx

.PHONEY: couchdb\:run redis\:start postgres\:start nginx\:start

