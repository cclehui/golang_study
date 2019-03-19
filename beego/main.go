package main

import (
	"fmt"
	"html/template"

	"babytree.com/test/beego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	beego.AddTemplateExt("tmpl")
	beego.AddViewPath("tmpl")

	beego.Router("test/template", &controllers.IndexController{}, "get:Template")

	beego.Get("index", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))

	})

	beego.Get("/template_raw", func(ctx *context.Context) {
		t, err := template.ParseFiles("tmpl/user/index.tmpl", "tmpl/user/header.tmpl", "tmpl/user/footer.tmpl")
		fmt.Println(33333333333333)
		//t, err := template.ParseFiles("tmpl/user/index.tmpl")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(t)

		err = t.ExecuteTemplate(ctx.ResponseWriter, "index.tmpl", &map[string]string{"title": "33333333333"})
		//err = t.ExecuteTemplate(c.Writer, "user/index.tmpl", &map[string]string{"title": "33333333333"})

		fmt.Println(err)
		//t.Execute(ctx.ResponseWriter, &map[string]string{"title": "222222222222"})
		//t.Execute(c.Writer, "aaaaaaaaaaaaaa")
	})

	beego.Run()

}
