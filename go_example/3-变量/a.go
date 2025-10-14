// 在 Go 中，变量 被显式声明，并被编译器所用来检查函数调用时的类型正确性

package main

import "fmt"

func main() {
	var a string = "initial" // var 声明 1 个或者多个变量。
	fmt.Println(a)

	var b, c int = 1, 2 // 你可以申明一次性声明多个变量。
	fmt.Println(b, c)

	var d = true // Go 将自动推断已经初始化的变量类型。
	fmt.Println(d)

	var e int // 声明变量而不给出对应的值将会赋予变量一个零值。
	fmt.Println(e)

	f := "apple" // := 语法是声明并初始化变量的简写方式。只能在函数体内使用。
	fmt.Println(f)
}
