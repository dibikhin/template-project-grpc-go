package server

type todoRepo struct {
	TODOs []string
}

func NewTODORepo() *todoRepo {
	return &todoRepo{TODOs: []string{}}
}
