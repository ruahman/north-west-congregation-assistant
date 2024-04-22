.PHONEY: up
up:
	@echo "Docker Compose Up"
	docker compose up -d 

.PHONEY: down
down:
ifdef v
	@echo "Docker Compose Down"
	docker compose down --volumes
else
	@echo "Docker Compose Down"
	docker compose down
endif

.PHONEY: start 
start:
	@echo "Docker Compose Start"
	docker compose start

.PHONEY: stop 
stop:
	@echo "Docker Compose Stop"
	docker compose stop

.PHONEY: bun
bun:
	@echo "Bun Shell"
	docker compose exec bun bash

.PHONEY: psql
psql:
	@echo "Postgres psql"
	docker compose exec postgres psql -U postgres

.PHONEY: postgres
postgres:
	@echo "Postgres shell"
	docker compose exec postgres bash


