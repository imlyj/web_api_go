package models

import (
//	"time"
//	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int64
	Name string
	Age int
	CreateTs int64
}

func (this *User) AddUser() error {
	if _, err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

func (this *User) All() []*User {
	var users []*User
	o.QueryTable("user").Limit(10000).All(&users)
	return users
}

func (this *User) Query() orm.QuerySeter {
	return o.QueryTable(this)
}