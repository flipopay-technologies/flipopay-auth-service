package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Application started")

	grpcServer := grpc.NewServer()

	log.Println("Grpc server started")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
