package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
//	"github.com/astaxie/beego"
)

type AppInfo struct {
	AppId  string  `orm:"pk;column(appId)"`
	BundleId  string `orm:"column(bundleId)"`
	Name  string
	Price float64
	Currency string
	Version string
	WebUrl string `orm:"column(webUrl)"`
	Ipad int
	Iphone int
	Update_date string
	Datas string
	Create_ts int64
	Update_ts int64
	Last_result int
	Is_enable int
	Weight  int
	MinimumOsVersion string   `orm:"column(minimumOsVersion)"`
	Category string
	Storage_group  string
	Is_paid_whitelist  int
	AppExtVrsId  string  `orm:"column(appExtVrsId)"`
}

func (this *AppInfo) TableName()string{
	return "AppInfo"
}

func (this *AppInfo) Add() error{
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

func (this *AppInfo) Get(fields ...string) error {
	if err := o.Read(this, fields...); err != nil {
//		if err == orm.ErrNoRows {
//			fmt.Println("------select no------")
//			return err
//		}
		return err
	}
	return nil
}

func (this *AppInfo) Update(fields ...string)error {
	if _,err := o.Update(this,fields ...);err != nil{
		return err
	}
	return nil
}

func (this *AppInfo) List() ([]AppInfo,error){
	var apps []AppInfo
	if _,err:=o.Raw("select * from AppInfo limit 10000").QueryRows(&apps);err!=nil{
		return apps,err
	}
	return apps,nil
}

func (this *AppInfo) ClearCategory(category string) error{
	if _,err := o.Raw("UPDATE AppInfo SET category = ?,weight = ? WHERE category = ?","",0,category).Exec();err!=nil{
		fmt.Println("clear category err",err)
		return err
	}
	return nil
}

//func (this *AppInfo) ListApp() []AppInfo{
//	var apps []AppInfo
//
//}

func (this *AppInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}


