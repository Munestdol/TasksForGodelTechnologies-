package service

import (
	"task5"
	"task5/pkg/repository"
)

type AdServices struct {
	repo repository.Ad
}

func NewAdService(repo repository.Ad) *AdServices{
	return &AdServices{repo: repo}
}

func (s *AdServices) Create(ad task5.Ad) (int,error){
	return s.repo.Create(ad)
}

func (s *AdServices) GetAllAds() ([]task5.Ad,error){
	return s.repo.GetAllAds()
}

func (s *AdServices) GetAdById(id int) (task5.Ad,error){
	return s.repo.GetAdById(id)
}

func (s *AdServices) GetAllAdsSortByPriceDesc() ([]task5.Ad,error){
	return s.repo.GetAllAdsSortByPriceDesc()
}

func (s *AdServices) GetAllAdsSortByPriceAsc() ([]task5.Ad,error){
	return s.repo.GetAllAdsSortByPriceAsc()
}

func (s *AdServices) GetAllAdsSortByDateDesc() ([]task5.Ad,error){
	return s.repo.GetAllAdsSortByDateDesc()
}

func (s *AdServices) GetAllAdsSortByDateAsc() ([]task5.Ad,error){
	return s.repo.GetAllAdsSortByDateAsc()
}

