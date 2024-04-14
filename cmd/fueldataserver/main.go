package main

import (
	"flag"
	"fmt"
	"log"
	"main/api/fueldata"
	"main/pkg/controllers"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 50051, "")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	fueldata.RegisterFuelDataServer(grpcServer, &controllers.FuelDataServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
