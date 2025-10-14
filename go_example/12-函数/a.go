// 函数 是 Go 的中心。我们将通过一些不同的例子来进行学习。

package main

import "fmt"

func plus(a int, b int) int { // 这里是一个函数，接受两个 int 并且以 int 返回它们的和
	return a + b // Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
}

func main() {
	res := plus(1, 2) // 正如你期望的那样，通过 name(args) 来调用一个函数，
	fmt.Println("1+2=", res)
}

// 这里有许多 Go 函数的其他特性。其中一个就是多值返回，也是我们接下来需要接触的。
