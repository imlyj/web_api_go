package utils

import (
	"fmt"
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
	u := Us{}
	u.Price = 1.0
	fmt.Println(u.Price)
}