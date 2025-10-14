// range 迭代各种各样的数据结构。让我们来看看如何在我们已经学过的数据结构上使用 rang 吧。

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums { // 这里我们使用 range 来统计一个 slice 的元素个数。数组也可以采用这种方法。
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // range 在 map 中迭代键值对
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" { // range 在字符串中迭代 unicode 编码。第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。
		fmt.Println(i, c)
	}

}
