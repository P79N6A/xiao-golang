// sync once
// author: baoqiang
// time: 2018/12/26 下午8:34
package lang

import (
	"sync"
	"code.byted.org/gopkg/pkg/log"
)

var one = sync.Once{}

func Once(){
	one.Do(func() {
		log.Printf("hahaha")
	})
}

func RunOnce(){

}