// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/rossifedericoe/go-dependecy-injection/app"
	"github.com/rossifedericoe/go-dependecy-injection/app/controllers"
	"github.com/rossifedericoe/go-dependecy-injection/app/services"
)

func InitializeApp() app.App {
	wire.Build(app.NewApp, controllers.NewBikeController, services.NewBikeService, controllers.NewCarController, services.NewCarService)
	return app.App{}
}
