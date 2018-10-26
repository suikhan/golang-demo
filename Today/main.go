package main

import (
	"./rests"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	rests.UserRest(app)

	app.Run(iris.Addr(":8080"))
}
