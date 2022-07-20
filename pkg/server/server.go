package server

import (
	"fmt"
	"log"
	"net"

	"TODO/pkg/api"
	"TODO/pkg/config"

	"google.golang.org/grpc"
)

func MakeServer() *grpc.Server {
	gr := NewTODORepo()
	s := NewTODOService(gr)
	gs := grpc.NewServer()

	api.RegisterTODOServer(gs, s)
	return gs
}

type TODORepo interface{}

func StartListen(cfg config.Config) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		log.Fatalf("server: %v", err)
	}
	return lis
}

func RunServer(srv *grpc.Server, lis net.Listener) {
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("server: serve: %v", err)
	}
	log.Println("server: stopped")
}
