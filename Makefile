install:
# @sudo apt update -y
	@sudo apt install protobuf-compiler -y
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0
	@export PATH="/home/alexkhoury/go/bin"

gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/main.proto

# Build the client executable
build:
	@go build -o bin/client client/main.go

# Run the client
run-client:
	@go run client/main.go

# Compiler et exécuter le fichier de transaction
run-transaction:
	@go run run_transactions/transaction.go

run-flood:
    go run flood_transactions.go


run-double-spend:
    go run double_spending/double_spending.go
# Clean build artifacts
clean:
	@rm -rf bin

# Default target
all: install gen build

# run:
# 	@go run server/main.go
# 	@go run client/main.go