package service

import (
	"github.com/uttam282005/go-backend-boilerplate/internal/lib/job"
	"github.com/uttam282005/go-backend-boilerplate/internal/repository"
	"github.com/uttam282005/go-backend-boilerplate/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
