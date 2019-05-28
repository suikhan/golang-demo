package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	v1 := app.Party("/api/v2", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/hello", func(ctx iris.Context) {
			ctx.WriteString("Hello from /hello")
		})
		v1.Get("/clinic", func(ctx iris.Context) {
			ctx.WriteString("Hello from /clinic")
		})
		v1.Post("/say", func(ctx iris.Context) {
			ctx.WriteString("say sent")
		})
		v1.Put("/say", func(ctx iris.Context) {
			ctx.WriteString("say updated")
		})
		v1.Delete("/say", func(ctx iris.Context) {
			ctx.WriteString("say deleted")
		})
	}

	app.Run(iris.Addr(":8081"))
}
