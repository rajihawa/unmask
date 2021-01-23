build:
	go build -o ./main main.go 

run:
	PORT=4000 ./main

test:
	go test ./tests