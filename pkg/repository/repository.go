package repository

import "github.com/jmoiron/sqlx"

type Authorisation interface {
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
	return &Repository{}
}
