run: build env
	@export `cat .env | xargs`; ./main

build: generate
	go build -o ./main cmd/main.go

generate:
	go generate ./...

test: build
	go test -v -race -covermode=atomic  -count=1  ./...

env: 
	cp .env.dist .env
