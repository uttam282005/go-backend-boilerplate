package repository

import "github.com/uttam282005/go-backend-boilerplate/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
