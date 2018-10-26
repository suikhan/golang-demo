package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/home", func(ctx iris.Context) {
			ctx.WriteString("Hello from /home")
		})
		v1.Get("/about", func(ctx iris.Context) {
			ctx.WriteString("Hello from /about")
		})
		v1.Post("/send", func(ctx iris.Context) {
			ctx.WriteString("sent")
		})
		v1.Put("/send", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/send", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
	}

	app.Run(iris.Addr(":8080"))
}
