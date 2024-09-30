package service

import (
	project "github.com/Chazarov/rest-app/entities"
	"github.com/Chazarov/rest-app/pkg/repository"
)

type AdvertsService struct {
	repo repository.Adverts
}

func NewAdvertsService(repo repository.Adverts) *AdvertsService {
	return &AdvertsService{repo: repo}
}

func (s *AdvertsService) Create(item project.AdvertItem) (int, error) {
	return s.repo.Create(item)
}
