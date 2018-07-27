# xiao-go-lang
&lt;go programming language> src code

1. 字符串格式化类型
- %d %x %o %b: 十进制整数，十六进制，八进制，二进制
- %f，%g, %e: 浮点数，自动浮点数，科学计数法
- %t: 布尔值
- %v: 内置格式的任何值
- %T: 任何值的类型
- %q: 带双引号的字符串或者带单引号的字符

2. map可以初始化为对应类型的初始值，而不需要类似python的defaultdict类似的东东。

3. 跟C语言不同，go只能使用i++，而且i++在go语言中等价于 i += 1

4. 跟C语言不同，switch,case语句不会默认的fallthrough

5. Go语言的关键字

6. | 关键词  | value                               |
   | ------ | ----------------------------------- |
   | 1      | break default func interface select |
   | 2      | case defer go map struct            |
   | 3      | chan else goto package switch       |
   | 4      | const fallthrough if range type     |
   | 5      | continue for import return var      |

   Go语言的预定义常量(&类型&函数)

   | 预定义常量(&类型&函数)    |                                           |
   | --------------------- | ----------------------------------------- |
   | 常量                   | true false iota nil                       |
   | 类型                   | int int8 int16 int32 int64                |
   | 类型                   | uint uint8 uint16 uint32 uint64 uintptr   |
   | 类型                   | float32 float64 complex128 complex64      |
   | 类型                   | bool byte rune string error               |
   | 函数                   | make len cap new append copy close delete |
   | 函数                   | complex real imag                         |
   | 函数                   | panic recover                             |

   ​
