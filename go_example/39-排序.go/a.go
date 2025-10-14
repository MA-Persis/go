// Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。我们首先关注内置数据类型的排序。

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"} // 排序方法是正对内置数据类型的；这里是一个字符串的例子。注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
