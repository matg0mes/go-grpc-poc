package main

import (
	"database/sql"
	"net"

	"github.com/matg0mes/go-grpc-poc/internal/database"
	"github.com/matg0mes/go-grpc-poc/internal/pb"
	"github.com/matg0mes/go-grpc-poc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db, err := sql.Open("slite3", "./db.sqlite")

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051") 
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
