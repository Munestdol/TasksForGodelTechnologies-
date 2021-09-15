package service

import (
	"task5"
	"task5/pkg/repository"
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

type Service struct {
	Ad
}

func NewServices(repos *repository.Repository) *Service{
	return &Service{
		Ad: NewAdService(repos.Ad),
	}
}
