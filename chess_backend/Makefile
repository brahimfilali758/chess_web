# Build target
build: 
	@go build -o ./bin/chess_backend
	@echo "Build complete"

run : build
	@./bin/chess_backend

test : 
	@go test -v ./...


.PHONY: build run test