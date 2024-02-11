.PHONEY: up
up:
	@echo "Docker Compose Up"
	docker compose up -d 

.PHONEY: start 
start:
	@echo "Docker Compose Start"
	docker compose start

.PHONEY: golang
golang:
	@echo "Golang Shell"
	dot_clean .
	docker compose exec golang bash

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

.PHONEY: down
down:
	@echo "Docker Compose Down"
	docker compose down --volumes

.PHONEY: stop 
stop:
	@echo "Docker Compose Stop"
	docker compose stop

.PHONEY: run 
run:
	@echo "...Run JW"
	go run main.go $(cmd)

.TEST: test 
test:
	@echo "...Testing"
	go test -v $(p)

.PHONEY: check
check:
	@echo "...Check"
	go build -o /dev/null .




