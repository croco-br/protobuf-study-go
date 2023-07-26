package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/croco-br/protobuf-study-go/internal/pb"
	"github.com/croco-br/protobuf-study-go/internal/pb/database"
	"github.com/croco-br/protobuf-study-go/internal/pb/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userDb := database.NewUser(db)
	userService := service.NewUserService(*userDb)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("running")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
