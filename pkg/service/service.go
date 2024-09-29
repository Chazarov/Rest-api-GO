package service

import (
	project "github.com/Chazarov/rest-app/entities"
	"github.com/Chazarov/rest-app/pkg/repository"
)

type Authorisation interface {
	CreateUser(user project.User) (int, error)
}

type Advertisement interface {
}

type AdvertisementList interface {
}

type Service struct {
	Authorisation
	Advertisement
	AdvertisementList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
	}
}
