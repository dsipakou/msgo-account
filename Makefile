.PHONY: migrate
migrate:
	migrate -path pkg/db/migrations -database 'postgres://postgres:postgres@localhost:5678/msgo_account?sslmode=disable' up

.PHONY: downgrade
downgrade:
	migrate -path pkg/db/migrations -database 'postgres://postgres:postgres@localhost:5678/msgo_account?sslmode=disable' down
