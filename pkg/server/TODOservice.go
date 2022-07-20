package server

import (
	"context"
	"log"

	"TODO/pkg/api"
)

type TODOService struct {
	repo TODORepo

	api.UnimplementedTODOServer
}

func NewTODOService(r TODORepo) *TODOService {
	return &TODOService{repo: r}
}

func (*TODOService) TODO(context.Context, *api.EmptyResponse) (*api.EmptyResponse, error) {
	log.Println("server: received TODO() call ok")
	return &api.EmptyResponse{}, nil
}
