package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteGzip([]byte("大噶好，我系过天了！"))
		ctx.Header("X-Custom", "Headers can be set here after WriteGzip as well, because the data are kept before sent to the client when using the context's GzipResponseWriter and ResponseRecorder.")
	})

	app.Get("/2", func(ctx iris.Context) {
		ctx.GzipResponseWriter().WriteString("大噶好，唔系渣渣辉！")
	})

	app.Run(iris.Addr(":8080"))
}
