package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: )" + p.name + "- age: " + strconv.Itoa(p.age) + "years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for _, element := range list {
		if value, ok := element.(int); ok {
			fmt.Println("int", value)
		} else if value, ok := element.(string); ok {
			fmt.Println("string", value)
		} else if value, ok := element.(Person); ok {
			fmt.Println("person", value)
		} else {
			fmt.Println("not known")
		}
	}

	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
	fmt.Println(p, v, x)

	q := reflect.TypeOf(x)
	fmt.Println(q)
}
