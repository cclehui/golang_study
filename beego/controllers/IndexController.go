package controllers

import (
	"fmt"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Template() {

	//this.Ctx.Output.Body([]byte("xddddddddddddddd"))

	this.Data["title"] = "this is IndexController, 1111111"

	fmt.Println(this.Data)

	//this.HtmlView("index.tmpl")
	this.HtmlView("user/index.tmpl")

}
