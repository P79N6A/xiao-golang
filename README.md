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

   | 关键词  | value                               |
   | ------ | ----------------------------------- |
   | 1      | break default func interface select |
   | 2      | case defer go map struct            |
   | 3      | chan else goto package switch       |
   | 4      | const fallthrough if range type     |
   | 5      | continue for import return var      |

6. Go语言的预定义常量(&类型&函数)

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

7. go语言中有零值保障机制，未初始化的变量会自动赋值为0

8. 调用new函数，返回的是对象的指针， var p *int = new(int)

9. 与c语言不同，函数里面返回变量的指针是安全的

10. 两个变量的类型不携带任何信息且是零值，例如struct{} 或者 []int，那么目前的实现里面，他们有相同的地址。针对new函数来讲。

11. 把函数里面的局部变量的地址赋值给全局变量，称作局部变量的逃逸，尽量避免这种情况，因为会阻碍垃圾回收。

12. map[], <-ch，x.(T),返回一个布尔型的变量，通常命名为ok，表示右边的操作有没有成功。

13. 类型通常会声明String()方法，表示该类型的人类可能形式，但是不要在里面使用fmt.Sprint等一系列相关的方法都会造成递归调用导致程序堆栈满异常。

14. Go语言目录：语言基础，表达式，函数，数据，方法，接口，并发，包，进阶

15. [极客学院Golang学习资料](http://wiki.jikexueyuan.com/project/the-go-study-notes-fourth-edition/language.html)

16. go lang keywords:
- break default func interface select
- case defer go map struct
- chan else goto package switch
- const fallthrough if range type
- continue for import return var

17. what's defer chan fallthrough?
- defer: finally run code
- chan: FIFO queue
- fallthrough: switch case fallthrough

18. predefined identifier:
- append bool type cap close complex complex64 complex128 uint16
- copy false float32 float64 imag int int8 int16 unit32
- int32 int64 iota len make new nil panic uint64
- print println real recover string true uint uint8 uintptr

19. type: rune(unicode)=int32 chan unitptr

20. 结构体，接口，函数，大小写区分可见性。大写可见

21. 常量结构体里面，如果后面一行没有初始化，那么就会使用上面一行的运算结果值

22. go语言中，数组是值类型，结构是值类型，切片是引用类型

23. copy(dest,src)

24. 当赋值类型到接口时，会发生拷贝，所以对原类型的修改，对接口无效
