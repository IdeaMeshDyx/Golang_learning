package functest

import (
	"fmt"
	"math"
	"testing"
)

func foo() {
	fmt.Println("foo")
}

// func with multiple return values
func bar(a int, b int) (int, int) {
	return a, b
}

// 回调函数
// 声明一个函数类型
type cb func(int) int

func testCallBack(x int, f cb) { //定义了一个函数 testCallBack
	f(x) //由于传进来的是callBack函数，该函数执行需要传入一个int类型参数，因此传入x
}

func callBack(x int) int {
	fmt.Printf("我是回调x%d\n", x)
	return x
}

// 匿名函数||函数闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// method：类似于类和对象
/* 定义结构体 */
type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func Test_main(t *testing.T) {
	foo()
	a, b := bar(1, 2)
	fmt.Println(a, b)

	// func as a parameter
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	// func call back
	testCallBack(1, callBack) //执行函数---testCallBack
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调x%d\n", x)
		return x
	})

	// 匿名函数||闭包
	number := getSequence()
	fmt.Println(number())
	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())

	// method
	c := Circle{3}
	fmt.Println(c.getArea())
}
