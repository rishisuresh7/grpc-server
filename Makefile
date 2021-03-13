build:
	@echo "Building binary..."
	@go build -o build/grpc-server main/main.go

run:
	@echo "Running grpc-server..."
	@go run main/main.go

clean:
	@echo "Cleaning build folder..."
	@rm -rf build

test:
	@echo "Running tests..."
	@go test ./...
