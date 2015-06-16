package api

import (
	"github.com/astaxie/beego"
	"web_api_go/models"
	"strings"
	"time"
)

type AppController struct {
	beego.Controller
}

func (this *AppController) AddApp(){
	var app models.AppInfo
	result := make(map[string]interface{})
	id := this.GetString("ids")
	if id == ""{
		result["rc"] = 1
		result["msg"] = "AppID列表为空"
		this.Data["json"] = result
		this.ServeJson()
	}
	ids := strings.Split(id,"\n")
	weight,_ := this.GetInt("weight")
	category := this.GetString("category")
	for _,appId := range ids{
		appId = strings.TrimSpace(appId)
		app.AppId = appId
		app.Create_ts = time.Now().Unix()
		app.Last_result = -1
		app.Is_enable = 1
		app.Weight =  weight
		app.Category = category
		app.Add()
	}
	result["rc"] = 1
	result["msg"] = "ok"
	this.Data["json"] = result
	this.ServeJson()
}

func (this *AppController)  ListApp(){
	var app models.AppInfo
	result := make(map[string]interface{})
	if apps ,err := app.List();err != nil{
		result["rc"] = 1
		result["msg"] = "get apps error"
	}else{
		result["rc"] = 0
		result["msg"] = "ok"
		result["data"] = apps
	}
	this.Data["json"] =  result
	this.ServeJson()
}

func  (this *AppController) UpdateCategory(){
	result := make(map[string]interface{})
	category := this.GetString("category","hot")
	appId := this.GetString("appId")
	weight,_ :=  this.GetInt("weigth",200)
	app := models.AppInfo{AppId:appId}
	if err := app.Get();err == nil{
		app.AppId = appId
		app.Weight = weight
		app.Category = category
		app.Update_ts = time.Now().Unix()
		app.Last_result = -1
		app.Is_enable = 1
		app.Add()
		result["rc"] = 0
		result["msg"] = "ok"
	}else{
		app.Update_ts = time.Now().Unix()
		app.Weight = weight
		app.Category = category
		app.Update()
		result["rc"] = 0
		result["msg"] = "ok"
	}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *AppController) ClearCategory(){
	var app models.AppInfo
	result := make(map[string]interface{})
	category := this.GetString("category")
	app.ClearCategory(category)
	result["rc"] = 0
	result["msg"] = "ok"
	this.Data["json"] = result
	this.ServeJson()
}

func (this *AppController) UpdateApp(){
	result := make(map[string]interface{})
	appId := this.GetString("appId")
	app := models.AppInfo{AppId:appId}
	if err := app.Get();err!= nil{
		result["rc"] = 1
		result["msg"] = "App is not in the AppInfo table."
	}else{
		app.Update_ts = time.Now().Unix()
		name := this.GetString("name")
		if name != ""{
			app.Name = name
		}
		price,_ := this.GetFloat("price")
		if price != 0{
			app.Price = price
		}
		currency := this.GetString("currency")
		if currency != ""{
			app.Currency = currency
		}
		version := this.GetString("version")
		if version != ""{
			app.Version = version
		}
		webUrl := this.GetString("webUrl")
		if webUrl != ""{
			app.WebUrl = webUrl
		}
		ipad,_ := this.GetInt("iPad")
		if ipad != 0{
			app.Ipad = ipad
		}
		iphone,_ := this.GetInt("iphone")
		if iphone != 0{
			app.Iphone = iphone
		}
		update_date := this.GetString("update_date")
		if update_date != ""{
			app.Update_date = update_date
		}
		datas := this.GetString("datas")
		if datas != ""{
			app.Datas = datas
		}
		category := this.GetString("category")
		if category != ""{
			app.Category = category
		}
		weight,_ := this.GetInt("weight")
		if weight != 0{
			app.Weight = weight
		}
		is_paid_whitelist,_ := this.GetInt("is_paid_whitelist")
		if is_paid_whitelist != 0{
			app.Is_paid_whitelist = is_paid_whitelist
		}
		appExtVrsId := this.GetString("appExtVrsId")
		if appExtVrsId != ""{
			app.AppExtVrsId = appExtVrsId
		}
		app.Update()
		result["rc"] = 0
		result["msg"] = "ok"
		this.Data["json"] = result
		this.ServeJson()
	}

}

func (this *AppController) UpdateWantApp(){
	result := make(map[string]interface{})
	appId := this.GetString("appId")
	if appId == ""{
		result["rc"] = 1
		result["msg"] = "AppID列表为空"
	}else{
		app := models.AppInfo{AppId:appId}
		if err := app.Get();err == nil{
			app.BundleId =  this.GetString("bundleId","")
			price,_ := this.GetFloat("price",0.0)
			app.Price = price
			app.Currency = this.GetString("currency","")
			app.Version = this.GetString("version","")
			app.Name = this.GetString("name","")
			ipad,_ := this.GetInt("iPad")
			app.Ipad = ipad
			iphone,_ := this.GetInt("iPhone")
			app.Iphone = iphone
			app.WebUrl = this.GetString("webUrl")
			app.Update_date = this.GetString("update_date","")
			app.MinimumOsVersion =  this.GetString("minimumOsVersion")
			app.Create_ts = time.Now().Unix()
			app.Update_ts =  time.Now().Unix()
			app.Is_enable = 1
			app.Last_result = -1
			app.Update()
			result["rc"] = 0
			result["msg"] = "ok"
		}
	}
	this.Data["json"] =result
	this.ServeJson()
}