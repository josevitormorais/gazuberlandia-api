migrate-create:
	migrate create -ext sql -dir db/migrations $(name)

migrate-up:
	migrate --path db/migrations -database "postgresql://docker:docker@localhost:5432/guberlandia?sslmode=disable" -verbose up

migrate-down:
	migrate --path db/migrations -database "postgresql://docker:docker@localhost:5432/guberlandia?sslmode=disable" -verbose down
