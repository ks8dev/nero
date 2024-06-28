build:
	go build -o ./cmd/. ./cmd/main.go

run:
	go run ./cmd/main.go

start: build
	./cmd/main