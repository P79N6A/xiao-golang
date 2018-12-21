// go语言
// author: baoqiang
// time: 2018/12/21 下午12:59
package lang

import "log"

func Info() {
	str := `
	1. 2007年9月开始设计，2009年11月正式发布
	2. 语言特点：编译，垃圾回收，并发	
	3. 应用领域：系统编程，web-server，存储架构
	`
	log.Printf("%s\n", str)
}
