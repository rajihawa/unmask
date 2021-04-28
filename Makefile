build:
	go build -o ./main main.go 

run:
	./main

test:
	go test ./tests/*


setup:
# install migrate CLI to create and apply migration files
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

gen-migration:
# example: make migrate name=create_example_table
	./migrate.linux-amd64 create -dir ./migrations -ext sql $(name)