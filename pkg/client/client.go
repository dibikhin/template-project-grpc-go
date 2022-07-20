package client

import (
	"log"

	"TODO/pkg/api"
	"TODO/pkg/config"

	"google.golang.org/grpc"
)

func Connect(cfg config.Config) (api.TODOClient, func()) {
	log.Println("client: connecting...")
	conn := grpcDial(cfg)
	client := api.NewTODOClient(conn)
	log.Println("client: connected")

	return client, func() {
		log.Println("client: disconnecting...")
		conn.Close()
		log.Println("client: disconnected")
	}
}

func grpcDial(cfg config.Config) *grpc.ClientConn {
	conn, err := grpc.Dial(
		cfg.Server.Address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.Server.Timeout),
	)
	if err != nil {
		log.Fatalf("client: grpc dial: %v", err)
	}
	return conn
}
