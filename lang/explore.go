// language explore
// author: baoqiang
// time: 2018/12/21 下午3:16
package lang

import "log"

func RunArray() {
	arr := [3]int{1, 2, 3}
	log.Printf("%T %+v\n", arr, arr)
}

//命名参数，需要注意函数内变量对其的修改
func Return() (a int) {
	//b := 2
	a = 1
	return
}

func Return2() (a int) {
	b := 2
	a = 1
	return b
}

func RunReturn() {
	log.Print(Return())
	log.Print(Return2())
}
