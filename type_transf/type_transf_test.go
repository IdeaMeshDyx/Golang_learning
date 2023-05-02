package typetransf

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_main(t *testing.T) {
	// 数值类型转换
	// 将整型转换为浮点型：
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)

	// string transform
	var str string = "10"
	var num int
	num, _ = strconv.Atoi(str)
	fmt.Println(num)

	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num2, str2)

	str3 := "3.14"
	num3, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str3, num3)
	}

	// interface transform
	/*
		接口类型转换有两种情况：类型断言和类型转换。
		类型断言用于将接口类型转换为指定类型，其语法为：
		value.(type)

		类型转换用于将指定类型型和接口类型互换，会在interface中细嗦，此处不赘述
	*/
	var i interface{} = "Hello, World"
	str4, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str4)
	} else {
		fmt.Println("conversion failed")
	}
}
