package main

import (
	"log"
	"net"

	pb "github.com/Hamiduzzaman96/Product---Service/Product---Service/proto/productpb"
	handler "github.com/Hamiduzzaman96/Product---Service/internal/handler/grpc"
	"github.com/Hamiduzzaman96/Product---Service/internal/repository"
	"github.com/Hamiduzzaman96/Product---Service/internal/usecase"
	"github.com/Hamiduzzaman96/Product---Service/pkg/database"
	"google.golang.org/grpc"
)

func main() {
	db := database.NewMySQl()

	repo := repository.NewProductRepository(db)

	uc := usecase.NewProductUsecase(repo)

	h := handler.NewProductHandler(uc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, h)

	log.Println("Product gRPC server running on :50051")
	server.Serve(lis)

}
