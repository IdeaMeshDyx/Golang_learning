package io_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Std_test() {
	// stdout / stdin and stderr
	for i := 0; i < 3; i++ {
		fmt.Println("hello")
	}
	print("guess where am i\n")
	for i := 0; i < 3; i++ {
		fmt.Println("hello")
	}
	/*
		fmt中的fmt.Println()是默认输出到stdout(standard output)的，而println是输出到stderr(standard error)，
		因此在IDE中看到的结果顺序是并不是预期的顺序
	*/
}

func Scanner_reader() {
	// scanner 封装了reader, 便于使用
	// 读取用户输入并打印
	fmt.Println("Scanner: 持续输入和输出, 输入exit退出")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "exit" {
			break
		} else {
			fmt.Println(str)
		}
	}

	fmt.Print("Reader: 输入内容, 按回车结束")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader.ReadString('\n')) // 以/n分割

}
func Scanner_matrix_test() {
	fmt.Println("read a matrix and print it")
	// 读取整数n
	fmt.Print("请输入一个整数n:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() //  默认以\n分割
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Print("请输入n*n矩阵:\n")
	matrix := make([][]int, n)
	for i := 0; i < n; i++ { // 输入n行, 每行n个整数, 以空格分割
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j], _ = strconv.Atoi(line[j])
		}
	}
	fmt.Printf("矩阵:%v", matrix)
}

func Reader_matrix_test() {
	fmt.Println("read a matrix and print it")
	fmt.Print("请输入一个整数n:")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\r\n")
	n, _ := strconv.Atoi(s)
	matrix := make([][]int, n)
	fmt.Println("请输入n*n矩阵:")
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\r\n")
		nums := strings.Split(line, " ")
		for j := 0; j < n; j++ {
			matrix[i][j], _ = strconv.Atoi(nums[j])
		}
	}
	fmt.Printf("矩阵:%v", matrix)
}

func fmt_scanner_test() {
	fmt.Println("input a matrix and print it")
	fmt.Println("请输入行数n和列数m:")
	var n, m int
	fmt.Scanf("%d%d\n", &n, &m)
	fmt.Println("请输入n行m列矩阵:")
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
		fmt.Scanf("\n")
	}
	fmt.Printf("矩阵:%v", matrix)
}

func Test_main(t *testing.T) {
	fmt.Println(`输入：	
				1: Std_test: test stdin and stdout and stderr
				2: Scanner_reader: test basic scanner and reader
				3: Scanner_matrix_test: test scanner
				4: Reader_matrix_test: test reader
				5: fmt_scanner_test: test fmt`)
	fmt.Println("\033[32m总结:windows下尽量用fmt.scanf(固定数量)和bufio.scanner(可变数量), Reader很麻烦\033[0m ")
	var num int
	fmt.Scanf("%d\n", &num) // 注意! 以\n分割
	switch num {
	case 1:
		Std_test()
	case 2:
		Scanner_reader()
	case 3:
		Scanner_matrix_test()
	case 4:
		Reader_matrix_test()
	case 5:
		fmt_scanner_test()
	default:
		fmt.Println("输入错误")
	}
}
