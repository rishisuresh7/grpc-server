build:
	@echo "Building binary..."
	@go build -o build/grpc-server main/main.go
	@echo "Built at build/grpc-server"

build-linux: clean
	@echo "Building binary for linux..."
	@GOOS=linux GOARCH=arm64 go build -ldflags "-X main.Version=1.0.0" -o build/grpc-server main/main.go
	@echo "Built at build/grpc-server"

run:
	@echo "Running grpc-server..."
	@source script.sh && go run main/main.go

clean:
	@echo "Cleaning build folder..."
	@rm -rf build
	@echo "Done"

test:
	@echo "Running tests..."
	@go test ./...
