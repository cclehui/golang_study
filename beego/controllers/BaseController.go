package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.ViewPath = "tmpl"
	this.TplExt = "tmpl"
}

func (this *BaseController) HtmlView(tmplName string) {

	this.TplName = tmplName

	res, err := this.RenderBytes()

	if err != nil {
		panic(err)
	}

	this.Ctx.Output.Body(res)

}
