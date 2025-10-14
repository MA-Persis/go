// Go 支持字符、字符串、布尔和数值 常量 。

package main

import (
	"fmt"
	"math"
)

const s string = "string" // const 用于声明一个常量。

func main() {
	fmt.Println(s)

	const n = 500000000 // const 语句可以出现在任何 var 语句可以出现的地方。
	fmt.Println(n)

	const d = 3e20 / n // 常量表达式可以进行任意精度的运算。
	fmt.Println(d)

	fmt.Println(int64(d)) // 数值常量没有固定的类型，直到被给定一个类型，比如通过显式转换。

	fmt.Println(math.Sin(n)) // 当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 math.Sin函数需要一个 float64 的参数。

	var s2 int = -5
	fmt.Println(math.Abs(float64(s2))) // 这里我们需要将 s 转换成一个 float64。

	fmt.Println(math.Max(3, 5.2))

	var f1 float64 = 3
	var f2 float64 = 5.2
	fmt.Println(math.Max(f1, f2))

	var f3 int = 3
	var f4 float64 = 5.2
	fmt.Println(math.Max(float64(f3), f4)) // 这里我们需要将 f3 转换成一个 float64。
}
