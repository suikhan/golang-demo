package main

import (
	"fmt"
	"io/ioutil"

	"../../CalcGZIP"
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
		json := CalcGZIP.DecodeJSON(encodeString)
		///对json字符串做一些处理工作

		encodeString = CalcGZIP.EncodeJSON(json)
		ctx.WriteString(encodeString)
	})

	app.Post("/ping", func(ctx iris.Context) {
		bytes, _ := ioutil.ReadAll(ctx.Request().Body) //读取请求数据
		encodeString := string(bytes)
		json := CalcGZIP.DecodeJSON(encodeString)
		///对json字符串做一些处理工作

		encodeString = CalcGZIP.EncodeJSON(json)
		ctx.WriteString(encodeString)
	})

	app.Put("/ping", func(ctx iris.Context) {
		bytes, _ := ioutil.ReadAll(ctx.Request().Body) //读取请求数据
		encodeString := string(bytes)
		json := CalcGZIP.DecodeJSON(encodeString)
		///对json字符串做一些处理工作

		encodeString = CalcGZIP.EncodeJSON(json)
		ctx.WriteString(encodeString)
	})

	app.Delete("/ping/{json:string}", func(ctx iris.Context) {
		encodeString := ctx.Params().GetString("json")
		json := CalcGZIP.DecodeJSON(encodeString)
		///对json字符串做一些处理工作

		encodeString = CalcGZIP.EncodeJSON(json)
		ctx.WriteString(encodeString)
	})

	es := CalcGZIP.EncodeJSON("abcdefghijk")
	fmt.Println("EncodeJSON:" + es)
	ms := CalcGZIP.DecodeJSON(es)
	fmt.Println("DecodeJSON:" + ms)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
