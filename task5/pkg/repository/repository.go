package repository

import (
	"github.com/jmoiron/sqlx"
	"task5"
)

type Ad interface {
	Create(ad task5.Ad) (int,error)
	GetAllAds() ([]task5.Ad,error)
	GetAllAdsSortByPriceDesc() ([]task5.Ad,error)
	GetAllAdsSortByPriceAsc() ([]task5.Ad,error)
	GetAllAdsSortByDateDesc() ([]task5.Ad,error)
	GetAllAdsSortByDateAsc() ([]task5.Ad,error)
	GetAdById(id int) (task5.Ad,error)
}

type Repository struct {
	Ad
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Ad: NewAdPostgres(db),
	}
}
