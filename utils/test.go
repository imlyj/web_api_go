package main

import (
	"fmt"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func pprint(v string){
	fmt.Println(v)
}


type Integer int

func (a *Integer) Add(b Integer){
	*a += b
}

func (a Integer) Ad(b Integer){
	a += b
}

type Us struct{
	Price float64
}

func main(){
//	fmt.Println(time.Now().Unix())
	a := time.Now().Unix()
	fmt.Println(time.Unix(a,0).Format("2006-01-02 15:04:05"))
}