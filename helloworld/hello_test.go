package hello

import (
	"fmt"
	"testing"
)

/*
demo: HelloWorld
写单元测试时, 测试函数应该以Test开头, 参数是t *testing.T, 文件要以_test结尾
*/

func Test_basic(t *testing.T) {
	fmt.Println("hello world")
}
