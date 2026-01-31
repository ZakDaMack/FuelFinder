package fuelfinder

import (
	"main/api/fuelfinder"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Commands   fuelfinder.FuelFinderClient
	Connection *grpc.ClientConn
}

// Creates a new FuelFinder Client to iteract with the server
func NewClient(target string) (*Client, error) {
	// set up grpc client
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// close conn once done, gen client
	// defer conn.Close()
	return &Client{
		Connection: conn,
		Commands:   fuelfinder.NewFuelFinderClient(conn),
	}, nil
}
