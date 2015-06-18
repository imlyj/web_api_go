package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
//	"github.com/astaxie/beego"
)

type AppInfo struct {
	AppId  string  `orm:"pk;column(appId)"json:"appId"`
	BundleId  string `orm:"column(bundleId)"json:"bundleId"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
	Currency string   `json:"currency"`
	Version string `json:"version"`
	WebUrl string `orm:"column(webUrl)" json:"webUrl"`
	Ipad int `json:"ipad"`
	Iphone int `json:"iphone"`
	Update_date string `json:"update_date"`
	Datas string `json:"datas"`
	Create_ts int64 `json:"create_ts"`
	Update_ts int64 `json:"update_ts"`
	Last_result int `json:"last_result"`
	Is_enable int  `json:"is_enable"`
	Weight  int  `json:"weigth"`
	MinimumOsVersion string   `orm:"column(minimumOsVersion)" json:"minimumOsVersion"`
	Category string `json:"category"`
	Storage_group  string `json:"storage_group"`
	Is_paid_whitelist  int  `json:"is_paid_whitelist"`
	AppExtVrsId  string  `orm:"column(appExtVrsId)" json:"appExtVrsId"`
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


