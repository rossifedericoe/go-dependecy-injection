package services

import (
	"github.com/rossifedericoe/go-dependecy-injection/app/domain"
)

type BaseService struct {

}

func NewBikeService() domain.BikeService {
	return &BaseService{}
}

func (service BaseService) GetColor() string {
	return "red"
}