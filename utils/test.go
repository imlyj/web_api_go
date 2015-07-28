package main

import (
	"fmt"
	"time"
	"runtime"
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

type Node struct{
	
	next interface{}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main(){
	fmt.Println(time.Now().Unix())
//	runtime.GOMAXPROCS(2)
	go say("world")
	say("hello")
}