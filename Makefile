# Makefile for simple Go API with profiling and benchmarks

APP_NAME := ./bin/simple-api
CMD_DIR := ./cmd/server

# Build the binary
build:
	go build -o $(APP_NAME) $(CMD_DIR)

# Run the server
run:
	go run $(CMD_DIR)

# Clean build artifacts
clean:
	rm -f $(APP_NAME)

# Run benchmarks
bench:
	go test ./... -bench=. -benchmem

# Generate CPU profile for 30 seconds
cpu-profile:
	go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30

# Generate heap profile
heap-profile:
	go tool pprof http://localhost:8080/debug/pprof/heap

.PHONY: build run clean bench cpu-profile heap-profile
