package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rossifedericoe/go-dependecy-injection/app/controllers"
)

type App struct {
	engine *gin.Engine
	homeController *controllers.HomeController
	carController *controllers.CarController
}

func NewApp(homeController *controllers.HomeController, carController *controllers.CarController) App {
	app := App{homeController:homeController,carController:carController}
	return app
}

func (app App) Run() {
	app.engine = gin.Default()
	app.engine.GET("/bike/color", app.homeController.Index)
	app.engine.GET("/car/km", app.carController.Index)
	app.engine.Run()
}
