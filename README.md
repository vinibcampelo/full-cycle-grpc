# Full Cycle gRPC Service

This repository contains a gRPC service built during the Full Cycle course.

## Tecnologies

- go  
- grpc  
- sqlite3

## How to run

To run the service, execute the following command in the terminal:

```bash
go mod tidy
go run cmd/grpcServer/main.go
```

## How to test

To test the service using a client, we recommend using the Evans tool. To install it, follow the instructions on the project page.

Once Evans is installed, run the following command in the terminal:

```bash
evans -r repl
```

This will start an interactive REPL (Read-Eval-Print Loop) of Evans, allowing you to send requests and receive responses from the gRPC service.

Remember that it's important to have a basic understanding of gRPC and the tools used in this project to use it effectively. Be sure to check out the gRPC and Evans documentation for more information.
