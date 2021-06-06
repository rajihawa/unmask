build:
	go build -o ./main main.go 

start:
	./main

run:
	go run

test:
	go test -tags testing ./...