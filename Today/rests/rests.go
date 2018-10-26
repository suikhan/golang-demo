package rests

import (
	"../dbs/gets"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func UserRest(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/user", getUsers)
	}
}

func getUsers(ctx iris.Context) {
	json, _ := gets.GetAllUsers()
	ctx.WriteString(json)
}
