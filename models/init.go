package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

func init(){
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_conf"))
	orm.RegisterModel(new(User),new(AppInfo))
	o = orm.NewOrm()
}