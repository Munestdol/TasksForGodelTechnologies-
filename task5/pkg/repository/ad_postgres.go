package repository

import (
	"task5"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostgres(db *sqlx.DB) *AdPostgres{
	return &AdPostgres{db: db}
}

func (r *AdPostgres) Create(ad task5.Ad) (int,error){
	var id int
	ad.Date = time.Now()

	query := fmt.Sprintf("INSERT INTO %s (title, price, main_photo, photo1, photo2, photo3, description, date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING id", adsTable)
	row := r.db.QueryRow(query, ad.Title, ad.Price, ad.MainPhoto,ad.Photo1,ad.Photo2,ad.Photo3, ad.Description, ad.Date)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AdPostgres) GetAllAds() ([]task5.Ad,error){
	var ads []task5.Ad
	query := fmt.Sprintf("SELECT * FROM %s" , adsTable)
	err := r.db.Select(&ads,query)

	return ads, err
}

func (r *AdPostgres) GetAllAdsSortByPriceDesc() ([]task5.Ad,error){
	var ads []task5.Ad
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY (price) DESC" , adsTable)
	err := r.db.Select(&ads,query)

	return ads, err
}

func (r *AdPostgres) GetAllAdsSortByPriceAsc() ([]task5.Ad,error){
	var ads []task5.Ad
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY (price) ASC" , adsTable)
	err := r.db.Select(&ads,query)

	return ads, err
}

func (r *AdPostgres) GetAllAdsSortByDateDesc() ([]task5.Ad,error){
	var ads []task5.Ad
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY (date) DESC" , adsTable)
	err := r.db.Select(&ads,query)

	return ads, err
}

func (r *AdPostgres) GetAllAdsSortByDateAsc() ([]task5.Ad,error){
	var ads []task5.Ad
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY (date) ASC" , adsTable)
	err := r.db.Select(&ads,query)

	return ads, err
}

func (r *AdPostgres) GetAdById(id int) (task5.Ad,error){
	var ad task5.Ad
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", adsTable)
	err := r.db.Get(&ad,query,id)

	return ad, err
}
