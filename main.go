package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/view"
)

func main() {
	app := iris.New()
	
	// load the ./templates/**.html
	templates := view.HTML("./templates", ".html"))
	app.AttachView(templates)
	
	app.Handle("GET", "/", func(ctx context.Context) {
	        // bind the {{ .Name }}
		ctx.ViewData("Name", "iris")
		// render the ./templates/hi.html
		ctx.View("hi.html")
	})

	app.Run(iris.Addr(":8080"))
}
