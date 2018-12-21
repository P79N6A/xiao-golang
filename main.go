// main方法
// author: baoqiang
// time: 2018/7/27 下午5:50
package main

import (
	"fmt"
	"time"
)

func run(){
	//sample.RunLissajous()
	//fmt.Println("hello")
	TestTime()
}

func tmp(){
	var p *int
	p = new(int)
	fmt.Println(p)
	fmt.Println(*p)
}

func TestTime(){
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	fmt.Println(tm2)
}

func main() {
	run()
	//tmp()
}
