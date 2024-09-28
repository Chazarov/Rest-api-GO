package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
