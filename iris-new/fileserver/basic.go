package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Favicon("./assets/favicon.ico")

	app.StaticWeb("/static", "./assets")

	app.Run(iris.Addr(":8080"))
}
