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
	4. 包注释使用/**/，注释后面空一行，注释使用大写开头
	`
	log.Printf("%s\n", str)
}

func Keywords() {
	keys := `
	1. fmt
	2. errors	
	3. strings
	4. strconv
	5. regexp
	6. encoding/json
	
	1. bool
	2. int,int8,int16,int32,int64
	3. uint,uint8(short),uint16,uint32,uint64
	4. float32,float64
	5. string
	6. complex64,complex128
	7. array

	1. slice
	2. map
	3. chan

	1. append
	2. close
	3. delete
	4. panic
	5. recover
	6. imag
	7. real
	8. make
	9. new
	10. cap
	11. copy
	12. len

	1. error
	type error interface {
		Error() string
	}

	`
	log.Printf("%s\n", keys)
}
