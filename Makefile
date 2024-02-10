.PHONEY: up
up:
	@echo "Docker Compose Up"
	docker compose up

.PHONEY: golang
golang:
	@echo "Golang Shell"
	docker compose exec golang bash

.PHONEY: bun
bun:
	@echo "Bun Shell"
	docker compose exec bun bash

.PHONEY: postgres
postgres:
	@echo "Postgres Shell"
	docker compose exec postgres psql -U postgres -h localhost

.PHONEY: down
down:
	@echo "Docker Compose Down"
	docker compose down
