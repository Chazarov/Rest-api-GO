package repository

import (
	project "github.com/Chazarov/rest-app/entities"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user project.User) (int, error)
	GetUser(username, password string) (project.User, error)
}

type Adverts interface {
	Create(item project.AdvertItem) (int, error)
}

type Repository struct {
	Authorisation
	Adverts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		Adverts:       NewAdwertsPostgres(db),
	}
}
