package ctlflowstmt

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	// if else
	var a int = 10
	var b int = 20
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	// switch and fallthrough, break
	// support multiple case
	var c int = 10
	c = 20
	switch c {
	case 10, 20:
		fmt.Println("10 or 20")
		fallthrough
	default:
		fmt.Println("default")

	}

	// type switch, can not fallthrough in type

	var d interface{} = 10
	// change d's type
	d = "hello"
	switch d.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}

	// select with channel, random select a channel with msg
	// 如果有 default 子句，则执行该语句。
	// 如果没有 default 子句，select 将阻塞，直到某个通道可以运行；Go 不会重新对 channel 或值进行求值。
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 10
	}()
	go func() {
		ch2 <- 20
	}()
	for i := 0; i < 2; i++ {
		select {
		case <-ch1:
			fmt.Println("ch1")
		case <-ch2:
			fmt.Println("ch2")
		}
	}

	// for
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print('\n')
	// for as while

	tempwhile := 0
	for tempwhile < 10 {
		fmt.Print(tempwhile)
		tempwhile++
	}
	fmt.Print('\n')

	// for range
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	for k, v := range m {
		fmt.Println(k, v)
	}

}
