package service

import "github.com/Chazarov/rest-app/pkg/repository"

type Authorisation interface {
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
	return &Service{}
}
