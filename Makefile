PHONY: build-cli

start-server:
	docker-compose up --build -d

build-cli:
	go build -o thumbnail-cli.exe ./cmd/cli/