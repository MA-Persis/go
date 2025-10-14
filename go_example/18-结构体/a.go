// Go 的结构体 是各个字段字段的类型的集合。这在组织数据时非常有用。

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) // 使用这个语法创建了一个新的结构体元素。

	fmt.Println(person{name: "Alice", age: 30}) // 你可以在初始化一个结构体元素时指定字段名字。

	fmt.Println(person{name: "Fred"}) // 省略的字段将被初始化为零值。

	fmt.Println(person{age: 20})

	fmt.Println(&person{name: "Ann", age: 40}) // & 前缀生成一个结构体指针

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // 使用点来访问结构体字段。

	s.age = 1 //结构体是可变的。
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age) // 也可以对结构体指针使用. - 指针会被自动解引用。

	sp.age = 51 //结构体是可变的。
	fmt.Println(sp.age)
}
