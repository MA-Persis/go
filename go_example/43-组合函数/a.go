// 我们经常需要程序在数据集上执行操作，比如选择满足给定条件的所有项，或者将所有的项通过一个自定义函数映射到一个新的集合上。

// 在某些语言中，会习惯使用泛型。Go 不支持泛型，在 Go 中，当你的程序或者数据类型需要时，通常是通过组合的方式来提供操作函数。

// 这是一些 strings 切片的组合函数示例。你可以使用这些例子来构建自己的函数。注意有时候，直接使用内联组合操作代码会更清晰，而不是创建并调用一个帮助函数。

package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int { // 返回目标字符串 t 出现的的第一个索引位置，或者在没有匹配值时返回 -1。
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool { // 如果目标字符串 t 在这个切片中则返回 true。
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool { // 如果这些切片中的字符串有一个满足条件 f 则返回true。
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool { // 如果切片中的所有字符串都满足条件 f 则返回 true。
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string { // 返回一个包含所有切片中满足条件 f 的字符串的新切片。
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string { // 返回一个对原始切片中所有字符串执行函数 f 后的新切片。
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
