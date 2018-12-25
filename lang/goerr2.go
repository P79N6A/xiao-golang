// go语言的错误处理机制
// author: baoqiang
// time: 2018/12/24 下午9:06
package lang

import (
	"errors"
	"fmt"
)

var ErrHttpConnect = errors.New("http connect err")
var ErrInvalidParam = fmt.Errorf("invalid param err")
