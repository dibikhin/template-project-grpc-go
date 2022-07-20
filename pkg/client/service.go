package client

import (
	"context"

	"TODO/pkg/api"
)

type todoService struct {
	c    api.TODOClient
	read func() string
}

func NewTODOService(c api.TODOClient, r func() string) *todoService {
	return &todoService{c, r}
}

func (s *todoService) TODO() error {
	_, err := s.c.TODO(context.TODO(), &api.EmptyResponse{})
	return err
}
