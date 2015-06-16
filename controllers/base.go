package controllers

import (
	"github.com/astaxie/beego"
	"web_api_go/utils"
)

type BaseRouter struct {
	beego.Controller
}

func (this *BaseRouter) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}