// Go 内建多返回值 支持。这个特性在 Go 语言中经常被用到，例如用来同时返回一个函数的结果和错误信息。

package main

import "fmt"

func vals() (int, int) { // (int, int) 在这个函数中标志着这个函数返回 2 个 int。
	return 3, 7
}

func main() {
	a, b := vals() // 这里我们通过多赋值 操作来使用这两个不同的返回值。
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // 如果你仅仅想返回值的一部分的话，你可以使用空白定义符 _。
	fmt.Println(c)
}
