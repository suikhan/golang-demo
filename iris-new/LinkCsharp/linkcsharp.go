package main

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/ping/{json:string}", func(ctx iris.Context) {
		encodeString := ctx.Params().GetString("json")
		bytes, _ := base64.StdEncoding.DecodeString(encodeString)
		str := string(bytes)
		ctx.WriteString("GET:" + str)
	})

	app.Post("/ping", func(ctx iris.Context) {
		bytes, _ := ioutil.ReadAll(ctx.Request().Body) //读取请求数据
		// encodeString := string(bytes)
		// bytes, _ = base64.StdEncoding.DecodeString(encodeString)
		str := string(bytes)
		ctx.JSON(str)
	})

	app.Put("/ping", func(ctx iris.Context) {
		bytes, _ := ioutil.ReadAll(ctx.Request().Body) //读取请求数据
		// encodeString := string(bytes)
		// bytes, _ = base64.StdEncoding.DecodeString(encodeString)
		str := string(bytes)
		ctx.JSON(str)
	})

	app.Delete("/ping/{json:string}", func(ctx iris.Context) {
		encodeString := ctx.Params().GetString("json")
		bytes, _ := base64.StdEncoding.DecodeString(encodeString)
		str := string(bytes)
		ctx.WriteString("DELETE:" + str)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
