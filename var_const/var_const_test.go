package varconst

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_main(t *testing.T) {
	// var	手动声明变量: var 变量名 变量类型 = 变量值
	var a string = "Test"
	fmt.Println(a)

	// :=	自动声明并初始化变量: 变量名 := 变量值
	b := 3
	fmt.Println(b)

	// 多个变量声明
	var c, d, e int
	c, d, e = 1, 2, 3
	fmt.Println(c, d, e)

	vname1, vname2, vname3 := "hello", "world", "!"
	fmt.Println(vname1, vname2, vname3)

	// 没有初始化时会自动赋值0
	var f int
	fmt.Println(f)

	/*
		当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
		可以通过 &i 来获取变量 i 的内存地址
		使用*来存这个地址，即指针
	*/
	var i int = 10
	var j int = i
	fmt.Printf("addr i:%x, addr j:%x\n", &i, &j)
	var k *int = &i
	fmt.Printf("pointer %x, content %d\n", k, *k)

	// 想要交换两个变量的值，则可以简单地使用 v1, v2 = v2, v1
	var v1, v2 int = 1, 2
	fmt.Println(v1, v2)
	v1, v2 = v2, v1
	fmt.Println(v1, v2)

	// const声明常量
	const pi float64 = 3.14
	fmt.Println(pi)

	//常量还可以用作枚举：

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown, Female, Male)

	// 计算表达式的大小可以用sizeof
	fmt.Println(unsafe.Sizeof(pi))

	/*
		字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
		计算字符串长度使用len
	*/
	s := "hello"
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(len(s))

	/*
		iota，特殊常量，enumeration constant generator,可以认为是一个可以被编译器修改的常量。
		当iota用于枚举常量时，它可以帮助您编写更简洁、更易于维护的代码。
	*/
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	// iota 的另类使用
	const (
		fst = 1 << iota
		sec = 3 << iota
		thd
		frth
	)
	fmt.Println(fst, sec, thd, frth)
	/*
		output: 0 6 12 24
		fst=1：左移 0 位，不变仍为 1。
		sec=3：左移 1 位，变为二进制 110，即 6。
		thd=3：左移 2 位，变为二进制 1100，即 12。
		frth=3：左移 3 位，变为二进制 11000，即 24。
	*/
}
