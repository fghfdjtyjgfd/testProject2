package core

import (
	"log"

	m "hexTest/model"
	repo "hexTest/repository"
)

// //service adabter////
type beerService struct {
	beerRepo repo.BeerRepository
}

func NewBeerService(beerRepo repo.BeerRepository) *beerService {
	return &beerService{beerRepo: beerRepo}
}

func (s *beerService) GetBeers(page int) ([]m.Beer, error) {
	beers, err := s.beerRepo.GetAll(page)
	if err != nil {
		log.Println(err)
	}

	return beers, nil
}

func (s *beerService) TotalPages() (int, error) {
	totalPages, err := s.TotalPages()
	if err != nil {
		log.Println(err)
	}
	return totalPages, nil
}



func (s *beerService) UpdateBeer(beer m.Beer) error {
	err := s.beerRepo.UpdateOne(beer)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) DeleteBeer(id int) error {
	err := s.beerRepo.DeleteOne(id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) CreateBeer(beer m.Beer) error {
	err := s.beerRepo.CreateAll(beer)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *beerService) CreateUser(user m.User) error {
	err := s.beerRepo.CreateUser(user)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *beerService) LoginUser(user m.User) (string, error) {
	t, err := s.beerRepo.LoginUser(user)
	if err != nil {
		log.Println(err)
	}
	return t, err
}
