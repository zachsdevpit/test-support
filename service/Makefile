build:
	go mod download && go mod verify
	go build -v -o bin/app

run: build
	@./bin/app
