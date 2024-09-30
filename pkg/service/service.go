package service

import (
	project "github.com/Chazarov/rest-app/entities"
	"github.com/Chazarov/rest-app/pkg/repository"
)

type Authorisation interface {
	CreateUser(user project.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Adverts interface {
	Create(item project.AdvertItem) (int, error)
}

type Service struct {
	Authorisation
	Adverts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		Adverts:       NewAdvertsService(repos.Adverts),
	}
}
