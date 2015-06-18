package main

import (
	_ "web_api_go/routers"
	"time"
	"github.com/astaxie/beego"

)

//自制模板函数
func int2str(in int64)string{
	return time.Unix(in, 0).Format("2006-01-02 15:04:05")
}

func main() {
	beego.AddFuncMap("int2str",int2str)
	beego.Run()
}

