package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rossifedericoe/go-dependecy-injection/app/domain"
)

type CarController struct {
	carService domain.CarService
}

func NewCarController(carService domain.CarService) *CarController {
	return &CarController{carService:carService}
}

func (controller CarController) Index(context *gin.Context) {
	context.JSON(200, fmt.Sprint(controller.carService.GetKM()))
}
