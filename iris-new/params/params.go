package main

import (
	"github.com/kataras/iris"
)

type LoginData struct {
	username string
	password string
}

func Login(ctx iris.Context) {
	var d LoginData

	err := ctx.ReadForm(&d)
	// err := ctx.ReadJSON(&d)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.WriteString("UserName=" + d.username + ", Password=" + d.password)
}

func main() {
	app := iris.New()

	app.Get("/login", Login)
	app.Post("/login", Login)

	app.Run(iris.Addr(":8080"))
}
