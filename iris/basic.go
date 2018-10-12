package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/basicauth"
	"time"
)

func newApp() *iris.Application {
	app := iris.New()

	authConfig := basicauth.Config{
		Users:   map[string]string{"suikhan": "suikhan", "admin": "admin"},
		Realm:   "Authorization Required",
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	app.Get("/", func(ctx context.Context) { ctx.Redirect("/admin") })

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)

		//http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)

		//http://localhost:8080/admin/setting
		needAuth.Get("/setting", h)
	}

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func h(ctx context.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s %s", ctx.Path(), username, password)
}
