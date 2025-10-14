// map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希 或者字典 ）。

package main

import "fmt"

func main() {
	m := make(map[string]int) // 要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type).

	m["k1"] = 7 // 使用典型的 make[key] = val 语法来设置键值对。
	m["k2"] = 13

	fmt.Println("map:", m) // 使用例如 Println 来打印一个 map 将会输出所有的键值对。

	v1 := m["k1"] // 使用 name[key] 来获取一个键的值
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m)) // 当对一个 map 调用内建的 len 时，返回的是键值对数目

	delete(m, "k2") // 内建的 delete 可以从一个 map 中移除键值对
	fmt.Println("map:", m)

	_, prs := m["k2"] // 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
