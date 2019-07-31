package main

import (
	"github.com/rossifedericoe/go-dependecy-injection/app/wire"
)

func main() {
	app := wire.InitializeApp()
	app.Run()
}
