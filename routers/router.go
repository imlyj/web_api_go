package routers

import (
	"web_api_go/controllers/api"
	"web_api_go/controllers/ui"
	"github.com/astaxie/beego"
)

func init() {
//	beego.Router("/", &controllers.MainController{})
//	beego.Router("/save",&controllers.DemoController{},"*:Save")
//	beego.Router("/test/all",&controllers.DemoController{},"*:All")

	beego.Router("/api/app/add",&api.AppController{},"*:AddApp")
	beego.Router("/api/app/list",&api.AppController{},"*:ListApp")
	beego.Router("/api/app/update",&api.AppController{},"*:UpdateApp")
	beego.Router("/api/app/update/want",&api.AppController{},"*:UpdateWantApp")
	beego.Router("/api/app/clear/category",&api.AppController{},"*:ClearCategory")
	beego.Router("/api/app/update/category",&api.AppController{},"*:UpdateCategory")
	//ui
	beego.Router("/ui/app/add",&ui.AppController{},"*:AddApp")
//	beego.Router("/ui/app/stat",&ui.AppController{},"*:Stat")
//	beego.Router("/ui/app/query",&ui.AppController{},"*:QueryApp")
	beego.Router("/ui/app/edit",&ui.AppController{},"*:EditApp")
	beego.Router("/ui/app/list",&ui.AppController{},"*:ListApp")



}
