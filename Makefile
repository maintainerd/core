run:
	go run cmd/server/main.go

migrate-up:
	./scripts/migrate.sh up

migrate-down:
	./scripts/migrate.sh down

migrate-down-all:
	./scripts/migrate.sh down-to 0

migrate-reset:
	./scripts/migrate.sh reset

seed-local:
	go run db/seeder/v1_seeder.go
