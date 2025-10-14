// for 是 Go 中唯一的循环结构。这里有 for 循环的三个基本使用方式。

package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // 最常用的方式，带单个循环条件。
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // 经典的初始化/条件/后置语句方式。
		fmt.Println(j)
	}

	for { // 无限循环。
		fmt.Println("loop")
		break
	}
}
