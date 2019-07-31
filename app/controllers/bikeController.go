package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rossifedericoe/go-dependecy-injection/app/domain"
)

type HomeController struct {
	baseService domain.BikeService
}

func NewBikeController(bikeService domain.BikeService) *HomeController {
	return &HomeController{baseService:bikeService}
}

func (controller HomeController) Index(context *gin.Context) {
	context.JSON(200, controller.baseService.GetColor())
}