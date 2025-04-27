.PHONY: build clean run build-linux build-darwin build-all

# Binary name
BINARY_NAME=s3-presigned-url-generator

# Build flags for production
LDFLAGS=-ldflags "-s -w"
GOFLAGS=-trimpath

# Build the application for current platform
build:
	go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_NAME) main.go

# Build for Linux AMD64
build-linux:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_NAME)-linux-amd64 main.go

# Build for macOS ARM64 (M1)
build-darwin:
	GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_NAME)-darwin-arm64 main.go

# Build for all platforms
build-all: build-linux build-darwin

# Clean build files
clean:
	go clean
	rm -f $(BINARY_NAME)*

# Run the application
run:
	go run main.go

# Build and run
build-run: build
	./$(BINARY_NAME)

# Install dependencies
deps:
	go mod tidy

# Update dependencies
update-deps:
	go get -u ./...
	go mod tidy 
