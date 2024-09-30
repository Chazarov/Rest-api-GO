package repository

import (
	"fmt"

	project "github.com/Chazarov/rest-app/entities"
	"github.com/jmoiron/sqlx"
)

type AdvertsPostgres struct {
	db *sqlx.DB
}

func NewAdwertsPostgres(db *sqlx.DB) *AdvertsPostgres {
	return &AdvertsPostgres{db: db}
}

func (r *AdvertsPostgres) Create(item project.AdvertItem) (int, error) {

	var id int
	createAdvertQuery := fmt.Sprintf("INSERT INTO %s (title, description, images_path, user_id) VALUES ($1, $2, $3, $4) RETURNING id", advertItemList)
	row := r.db.QueryRow(createAdvertQuery, item.Title, item.Description, item.ImagesPath, item.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
