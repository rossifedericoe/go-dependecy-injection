package services

import "github.com/rossifedericoe/go-dependecy-injection/app/domain"

type CarService struct {
}

func NewCarService() domain.CarService {
	return &CarService{}
}

func (service *CarService) GetKM() int {
	return 0
}
