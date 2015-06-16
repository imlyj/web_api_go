package api

import (
	"github.com/astaxie/beego"
	"web_api_go/models"
//	"encoding/json"
//	"fmt"
//	"github.com/astaxie/beego/orm"
)

type DemoController struct {
	beego.Controller
}

var user models.User

var users  []*models.User

func (this *DemoController) Save(){
	name := this.GetString("name")
	age,_ := this.GetInt("age")
	user.Age  = age
	user.Name = name
	user.AddUser()
	this.Ctx.WriteString("save ok")
}

func (this *DemoController) All(){
	result := make(map[string]interface{})
	result["rc"] = 0
	result["msg"] = "ok123"
	result["data"] = user.All()
	this.Data["json"] = result
	this.ServeJson()
//	if value,err := json.Marshal(result);err != nil{
//		fmt.Println("json err",err)
//	}else{
//		this.Ctx.WriteString(string(value))
//	}
}