package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/s", suikhan)

	app.Run(iris.Addr(":8080"))
}

func suikhan(ctx iris.Context) {
	ctx.ViewData("message", "suikhan")
	ctx.View("hello.html")
}
