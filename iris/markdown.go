package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/context"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	app := iris.New()

	app.Get("/{fileName}", cache.Handler(10*time.Second), writeMarkdown)
	app.Run(iris.Addr(":8080"))
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func writeMarkdown(ctx context.Context) {
	fileName := ctx.Params().Get("fileName")
	cc, _ := ReadAll(fileName)
	println("Handler executed. Content refreshed.")
	ctx.Markdown(cc)
}
