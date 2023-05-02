package arrayslice

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	// array
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[1])

	//  将索引为 1 和 3 的元素初始化
	balance := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance)

	// multi-dimensional array
	var b [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(b)

	//使用 append() 函数向空的二维数组添加两行一维数组
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	fmt.Println(values)

	// slice
	var c []int = []int{1, 2, 3}
	fmt.Println(c)
	c = append(c, 4)
	fmt.Println(c)

	var slice1 []int = make([]int, 3, 4)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	fmt.Println(slice1)

	// slice's slice. left is closed field, right is open field
	var slice2 []int = slice1[2:4]
	fmt.Println(slice2)

	/*
		我们基于原数组或者切片创建一个新的切片后，那么新的切片的大小和容量是多少呢？
		这里有个公式，对于底层数组容量是 k 的切片 slice[i:j] 来说：
		长度: j-i
		容量: k-i
	*/
	fmt.Println(cap(slice2))

	// 切片是可索引的，并且可以由 len() 方法获取长度。
	// 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
	// 如果切片的容量小于1024个元素，那么扩容的时候slice的cap就乘以2；一旦元素个数超过1024个元素，则乘以1.25
	var slice3 []int = make([]int, 3)
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 1)
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 2)
	fmt.Println(cap(slice3))

	// array、slice、map、channel是可迭代的
	// 这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	// 第一个表示下标（key）, 第二个表示值
	nums := []int{2, 3, 4}
	sum := 0
	// use _ to ignore the value
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}
