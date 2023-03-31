package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/devfullcycle/14-grpc/internal/database"
	"github.com/devfullcycle/14-grpc/internal/pb"
	"github.com/devfullcycle/14-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id TEXT PRIMARY KEY, name TEXT, description TEXT)")
	if err != nil {
		log.Fatalf("Error creating categories table: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
