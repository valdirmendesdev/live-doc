package services

import "github.com/valdirmendesdev/live-doc/internal/core/repositories"

type Service struct {
	repository repositories.UserRepository
}

func NewCustomer(r repositories.UserRepository) *Service {
	return &Service{
		repository: r,
	}
}
