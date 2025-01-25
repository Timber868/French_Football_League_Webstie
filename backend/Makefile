build:
	@go build -o bin/french-league cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/french-league

#need to install migrate via go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#make migration add-user-table
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations/ $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down