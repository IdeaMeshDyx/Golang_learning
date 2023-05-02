package interfacetest

import (
	"encoding/binary"
	"fmt"
	"testing"
)

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

type IntWriter struct {
	num int
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func (iw *IntWriter) Write(data []byte) (int, error) {
	iw.num += int(data[0])
	return 1, nil
}

func Test_main(t *testing.T) {
	/*
		如上，接口类型是一组函数的集合，由结构体来实现
		Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
		接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
		Go 语言中的接口是隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口。因此，我们可以通过将接口作为参数来实现对不同类型的调用，从而实现多态。

		可以类比为C++的类和对象，接口就是虚函数，结构体就是子类
		结构体实现了这个接口，那这个接口方法就成为这个字类的成员函数
	*/

	// 同时，接口可以类型转换
	var w Writer = &StringWriter{}
	var w2 Writer = &IntWriter{}

	sw := w.(*StringWriter)
	sw.Write([]byte("hello"))

	var n uint16 = 65535
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, n)
	iw := w2.(*IntWriter)
	iw.Write(data)

	var m uint8 = 255
	iw.Write([]byte{m})

	fmt.Println(sw.str)
	fmt.Println(iw.num)
}
