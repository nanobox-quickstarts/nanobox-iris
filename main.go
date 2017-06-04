package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.Writef("hello world\n")
	})

	app.Run(iris.Addr(":8080"))
}
