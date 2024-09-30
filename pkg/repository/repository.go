package repository

import (
	project "github.com/Chazarov/rest-app/entities"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user project.User) (int, error)
	GetUser(username, password string) (project.User, error)
}

type Advertisement interface {
}

type AdvertisementList interface {
}

type Repository struct {
	Authorisation
	Advertisement
	AdvertisementList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
