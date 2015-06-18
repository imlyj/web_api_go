package ui

import (
	"github.com/astaxie/beego"
	"web_api_go/models"
	"web_api_go/utils"
	"strings"
//	"time"
//	"fmt"
)

type AppController struct {
	beego.Controller
}

func (this *AppController) AddApp(){
	this.TplNames = "ui_add_app.html"
}

func (this *AppController) EditApp(){
	result :=  make(map[string]interface{})
	appId := this.GetString("appId")
	if appId == ""{
		result["rc"] = 1
		result["msg"] = "appId不能为空"
	}else{
		app := models.AppInfo{AppId:appId}
		if err := app.Get();err != nil{
			result["rc"] = 1
			result["msg"] = "无数据"
		}else{
			result["rc"] = 0
			result["msg"] = "ok"
			result["data"] = app
		}
	}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *AppController) ListApp(){
	result :=  make(map[string]interface{})
	var page int64
	var pagesize int64 = 20
	var list []*models.AppInfo
	var app models.AppInfo
	if page, _ = this.GetInt64("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize
	query := app.Query()
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-update_ts").Limit(pagesize, offset).All(&list)
		result["rc"] = 0
		result["msg"] = "ok"
		result["data"] = list
		result["objPage"] = models.NewObjPage(count,page,pagesize)
	}else{
		result["rc"] = 1
		result["msg"] = "没有你要查询的数据！"
	}
	this.Data["json"] = result
	this.ServeJson()
//	this.Data["list"] = list
//	this.Data["pagebar"] = models.NewPager(page, count, pagesize, "/ui/app/list?page=%d").ToString()
//	this.TplNames = "app_list.html"
}

func (this *AppController) QueryApp(){
	result :=  make(map[string]interface{})
	var app models.AppInfo
	var list []*models.AppInfo
	var startTimeInt int64
	var endTimeInt int64
	appId := strings.TrimSpace(this.GetString("appId",""))
	startTime := strings.TrimSpace(this.GetString("start_time",""))
	endTime := strings.TrimSpace(this.GetString("end_time",""))
	page,_ := this.GetInt64("page",1)
	pagesize,_ := this.GetInt64("page_size",20)
	name := strings.TrimSpace(this.GetString("name",""))

	if startTime !="" && endTime !="" {
		startTimeInt = utils.Str2int(startTime)
		endTimeInt = utils.Str2int(endTime)
	}
	offset := (page-1)*pagesize
	query := app.Query()
	count ,_ := query.Filter("appId",appId).Filter("update_ts__gte",startTimeInt).Filter("update_ts__lte",endTimeInt).Filter("name__iexact",name).Limit(pagesize, offset).Count()
	if count > 0 {
		query.Filter("appId",appId).Filter("update_ts__gte",startTimeInt).Filter("update_ts__lte",endTimeInt).Filter("name__iexact",name).Limit(pagesize, offset).All(&list)
		result["rc"] = 0
		result["msg"] = "ok"
		result["data"] = list
		result["objPage"] = models.NewObjPage(count,page,pagesize)
	}else {
		result["rc"] = 1
		result["msg"] = "没有你要查询的数据！"
	}
	this.Data["json"] = result
	this.ServeJson()
//	this.Data["list"] = list
//	this.Data["pagebar"] = models.NewPager(page, count, pagesize, "/ui/app/list?page=%d").ToString()
//	this.TplNames = "app_list.html"
}

